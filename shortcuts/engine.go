package shortcuts

import "github.com/abolfazlbeh/zhycan/internal/engine"

// RegisterRestfulApp - register the restful application to the engine
func RegisterRestfulApp(app engine.RestfulApp) error {
	err := app.Routes()
	if err != nil {
		return err
	}
	return nil
}
