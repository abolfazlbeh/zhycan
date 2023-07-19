package shortcuts

//
//import (
//	"github.com/spf13/cobra"
//)
//
//// Define Types
//
////type Command struct {
////	// Use is the one-line usage message.
////	// Recommended syntax is as follows:
////	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
////	//   ... indicates that you can specify multiple values for the previous argument.
////	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
////	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
////	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
////	//       optional, they are enclosed in brackets ([ ]).
////	// Example: add [-F file | -D dir]... [-f format] profile
////	Use string
////
////	// Aliases is an array of aliases that can be used instead of the first word in Use.
////	Aliases []string
////
////	// Short is the short description shown in the 'help' output.
////	Short string
////
////	// Long is the long message shown in the 'help <this-command>' output.
////	Long string
////
////	// Example is examples of how to use the command.
////	Example string
////
////	// Deprecated defines, if this command is deprecated and should print this string when used.
////	Deprecated string
////
////	// Version defines the version for this command. If this value is non-empty and the command does not
////	// define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
////	// will print content of the "Version" variable. A shorthand "v" flag will also be added if the
////	// command does not define one.
////	Version string
////
////	// Run: Typically the actual work function. Most commands will only implement this.
////	Run func(cmd *Command, args []string)
////
////	// Hidden defines, if this command is hidden and should NOT show up in the list of available commands.
////	Hidden bool
////
////	// cobra command instance
////	cobraCommand *cobra.Command
////}
//
////type CompletionOptions cobra.CompletionOptions
//
//// Group - Structure to manage groups for commands
//type Group struct {
//	cobra.Group
//}
//
//// Command - is just that, a command for your application.
//// E.g.  'go run ...' - 'run' is the command. Cobra requires
//// you to define the usage and description as part of your command
//// definition to ensure usability.
//type Command struct {
//	cobra.Command
//}
//
//// CompletionOptions - are the options to control shell completion
//type CompletionOptions struct {
//	cobra.CompletionOptions
//}
//
////func (c *Command) Execute() error {
////	if c.cobraCommand == nil {
////	}
////
////	c.cobraCommand.Run = func(cmd *cobra.Command, args []string) {
////		c.Run(c, args)
////	}
////
////	return c.cobraCommand.Execute()
////}
//
//func tttt(dc string) {
//	c := &Command{}
//	c.CompletionOptions = CompletionOptions{}.CompletionOptions
//	c.Run
//}
