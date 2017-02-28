package cbcluster

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

/*

   #import cb_bootstrap

   # Wrapper around bootstrap.couchbase.io REST API service which has global view of
   # cluster and can track which node is the boostrap

   #couchbase_cluster = cb_bootstrap.CouchbaseCluster(cluster_token, node_id)
   #couchbase_cluster.SetAdminUser("Administrator")
   #couchbase_cluster.SetAdminPassword("Password")
   #couchbase_cluster.SetCouchbaseServerName(socket.gethostname())  # how to get the public ip?
   #couchbase_cluster.WireUp()  # blocks until it either sets up as initial node or joins other nodes
   #couchbase_cluster.AddBucketIfMissing(
   #   Name="data-bucket",
   #   PercentRam=0.50,
   #)
   #couchbase_cluster.AddBucketIfMissing(
   #   Name="index-bucket",
   #   PercentRam=0.50,
   #)
*/

type CouchbaseCluster struct {
	ClusterId string                    // Something to uniquely identify the cluster
	DynamoDb  dynamodbiface.DynamoDBAPI // DynamoDB driver or mock
}

type CouchbaseNode struct {
	CouchbaseCluster *CouchbaseCluster
	IpAddrOrHostname string // The ip address or hostname for this Couchbase Node
	InitialNode      bool   // Whether this is the initial node that others can join
}

func (cluster *CouchbaseCluster) CreateOrJoinCuster(iPAddrOrHostname string) (CouchbaseNode, error) {


	// Create a new cluster object from database, or retrieve existing
	// -------------------------------------------------------------------------------------------------------------
	putItemInput := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"cluster_id": {
				S: aws.String(cluster.ClusterId),
			},
			"initial_node_ip_addr_or_hostname": {
				S: aws.String(iPAddrOrHostname),
			},
		},
		TableName:           aws.String("cb-bootstrap"),
		ConditionExpression: aws.String("attribute_not_exists(cluster_id)"),
	}
	putItemOutPut, err := cluster.DynamoDb.PutItem(putItemInput)
	if err == nil {

		// Create succeeded -- which means iPAddrOrHostname successfully became the inital node

		cbNode := cluster.NewCouchbaseNode()

		// no error,
		cbNode.InitialNode = true
		cbNode.IpAddrOrHostname = iPAddrOrHostname

		return cbNode, nil
	}

	// Create failed -- if due to existing cluster, then fetch existing cluster details, or else raise error
	// -------------------------------------------------------------------------------------------------------------

	// We got an error.  If it was just a ErrCodeConditionalCheckFailedException,
	// then we should just do a GetItem call to get the value
	awsErr, ok := err.(awserr.Error)
	if !ok {
		// unexpected error
		log.Printf("Expected an awserr.Error, got: %+v", err)
		return CouchbaseNode{}, err
	}

	if awsErr.Code() != dynamodb.ErrCodeConditionalCheckFailedException {
		// unexpected error
		log.Printf("Expected an awserr.Error with dynamodb.ErrCodeConditionalCheckFailedException, got :%+v", awsErr)
		return CouchbaseNode{}, err
	}
	
	log.Printf("Cluster already exists!  Fetching existing initial node from db.  PutItemOutput: %+v", err, putItemOutPut)

	// now we need to do a fetch to get the initial node ip addr or host
	cbNode := cluster.NewCouchbaseNode()

	err = cbNode.LoadFromDatabase()
	if err != nil {
		return CouchbaseNode{}, err
	}

	log.Printf("Loaded cbnode from db: %+v", cbNode)

	cbNode.InitialNode = false

	return cbNode, nil

}

func (cluster *CouchbaseCluster) NewCouchbaseNode() CouchbaseNode {

	return CouchbaseNode{
		CouchbaseCluster: cluster,
	}

}


func (cbnode *CouchbaseNode) LoadFromDatabase() error {

	attribute := dynamodb.AttributeValue{S: aws.String(cbnode.CouchbaseCluster.ClusterId)}
	query := map[string]*dynamodb.AttributeValue{"cluster_id": &attribute}

	getItemInput := &dynamodb.GetItemInput{
		Key: query,
		ConsistentRead: aws.Bool(true),
		TableName:           aws.String("cb-bootstrap"),
	}
	getItemOutput, err := cbnode.CouchbaseCluster.DynamoDb.GetItem(getItemInput)
	if err != nil {
		return err
	}

	initialNodeIpOrHostnameAttribute := getItemOutput.Item["initial_node_ip_addr_or_hostname"]
	cbnode.IpAddrOrHostname = *initialNodeIpOrHostnameAttribute.S

	return nil

}


func CreateDynamoDbSession() *dynamodb.DynamoDB {
	// connect to dynamodb
	awsSession := session.New()
	dynamoDb := dynamodb.New(awsSession)
	return dynamoDb
}
