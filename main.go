package main

import (
	"cdk.tf/go/stack/generated/aws"
	"github.com/aws/jsii-runtime-go"
	"cdk.tf/go/stack/generated/aws/ec2"
	"cdk.tf/go/stack/generated/aws/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("us-east-1"),
	})

	sg := vpc.NewSecurityGroup(stack, jsii.String("security-group"), &vpc.SecurityGroupConfig{
		Name:        jsii.String("cdktf-typeScript-demo-sg"),
		Description: jsii.String("allow traffic to the ec2 instance"),
		Ingress: []vpc.SecurityGroupIngress{
			vpc.SecurityGroupIngress{
				Protocol:   jsii.String("tcp"),
				FromPort:   jsii.Number(80),
				ToPort:     jsii.Number(80),
				CidrBlocks: &[]*string{jsii.String("0.0.0.0/0")},
			},
			vpc.SecurityGroupIngress{
				Protocol:   jsii.String("tcp"),
				FromPort:   jsii.Number(22),
				ToPort:     jsii.Number(22),
				CidrBlocks: &[]*string{jsii.String("0.0.0.0/0")},
			},
			vpc.SecurityGroupIngress{
				Protocol:   jsii.String("tcp"),
				FromPort:   jsii.Number(443),
				ToPort:     jsii.Number(443),
				CidrBlocks: &[]*string{jsii.String("0.0.0.0/0")},
			},
		},
		Egress: []vpc.SecurityGroupEgress{
			vpc.SecurityGroupEgress{
				Protocol: jsii.String("-1"),
				FromPort: jsii.Number(0),
				ToPort:   jsii.Number(0),
			},
		},

		Tags: &map[string]*string{
			"Name":    jsii.String("Security-Group-Golang-Ec2"),
		},
	})

	ec2.NewInstance(stack, jsii.String("ec2-instance"), &ec2.InstanceConfig{
		Ami:                 jsii.String("ami-22d399vb13b7d22c2"),
		InstanceType:        jsii.String("t3.medium"),
		KeyName:             jsii.String("rookey"),
		VpcSecurityGroupIds: &[]*string{sg.Id()},
		Tags: &map[string]*string{
			"Name":    jsii.String("cdk-ec2-instance"),
		},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	stack := NewMyStack(app, "cdktf-go-ec2")
	cdktf.NewRemoteBackend(stack, &cdktf.RemoteBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("xxxxxxxxxx"),
		Workspaces:   cdktf.NewNamedRemoteWorkspace(jsii.String("cdktf-ec2")),
	})

	app.Synth()
}
