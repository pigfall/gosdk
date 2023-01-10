package mongo

import(
	"fmt"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func ObjectIDCommand()*cobra.Command{
	cmd := &cobra.Command{
		Use: "objectID",
	}

	objectIDCreateCmd := &ObjectIDCreateCommand{}
	objectIDCreateCommand := &cobra.Command{
		Use: "create",
		RunE: objectIDCreateCmd.Run,
	}


	cmd.AddCommand(
		objectIDCreateCommand,
	)

	return cmd
}

type ObjectIDCreateCommand struct{}

func (c *ObjectIDCreateCommand) Run(cmd *cobra.Command,args []string)error{
	id := primitive.NewObjectID().Hex()
	fmt.Println(id)

	return nil
}
