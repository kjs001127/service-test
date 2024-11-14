package config

import (
	"embed"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/viper"

	"github.com/channel-io/ch-app-store/internal/util"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/ddb"
	chlog "github.com/channel-io/ch-app-store/lib/log/channel"
)

type Stage string

const (
	EXP        = Stage("exp")
	PRODUCTION = Stage("production")
	FEATURE    = Stage("feature")
	TEST       = Stage("test")
	DEV        = Stage("development")
	CI         = Stage("ci")
)

var (
	cfg          *Config
	cfgExtension = "yaml"
	validStages  = []Stage{EXP, PRODUCTION, FEATURE, TEST, DEV, CI}
)

//
//go:embed *
var configFiles embed.FS

type Config struct {
	Stage              string `required:"true" name:"config.stage"`
	ServiceName        string `required:"true" name:"config.serviceName"`
	ChannelServiceName string `required:"true"`
	Port               string `required:"true"`
	Auth               struct {
		AuthAdminURL  string `required:"true"`
		JWTServiceKey string `required:"true"`
	}
	Log                chlog.Config
	Psql               db.Config
	DDB                ddb.DDBConfig
	SystemLogTableName string `required:"true"`
	RateLimiter        struct {
		URL           string `required:"true"`
		AppRuleID     string `required:"true"`
		ManagerRuleID string `required:"true"`
		UserRuleID    string `required:"true"`
	}
	Services util.ServiceMap `required:"true"`
}

func init() {
	fillDefaultValues()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func fillDefaultValues() {
	viper.SetDefault("stage", "development")
	viper.SetDefault("port", "3020")
}

func Get() *Config {
	if cfg == nil {
		cfg = &Config{}
		loadConfig(cfg)
		validate(cfg)
	}
	return cfg
}

func loadConfig(cfg *Config) {
	// Set config from jsom file
	viper.SetConfigType(cfgExtension)

	cfgStr := configAsString()
	c := replaceEnvVars(cfgStr)
	if err := viper.ReadConfig(strings.NewReader(c)); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err)
	}
}

func configAsString() string {
	path := fmt.Sprintf("%s.%s", currentStage(), cfgExtension)
	fmt.Printf("using config file %s\n", path)
	bytes, err := configFiles.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("config not found: %v", err))
	}
	return string(bytes)
}

func currentStage() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	return env
}

// replaceEnvVars replaces ${ANY_ENV_VAR} in config file to the environment variable ANY_ENV_VAR
func replaceEnvVars(c string) string {
	re := regexp.MustCompile(`\${[^}]+}`)
	return re.ReplaceAllStringFunc(c, func(matched string) string {
		envName := matched[2 : len(matched)-1]
		val := os.Getenv(envName)
		if _, err := strconv.Atoi(val); err == nil {
			return val
		}
		if _, err := strconv.ParseBool(val); err == nil {
			return val
		}
		return fmt.Sprintf("\"%s\"", val)
	})
}

// / Validation
func validate(c *Config) {
	if err := validateStage(c); err != nil {
		panic(err)
	}

	if err := validateStruct(c); err != nil {
		panic(err)
	}
}

func validateStage(cfg *Config) error {
	for _, validStage := range validStages {
		if validStage == Stage(cfg.Stage) {
			return nil
		}
	}
	return fmt.Errorf("invalid stage, stage = %s", cfg.Stage)
}

func validateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				// Create empty struct to check fields in nil pointer struct
				field = reflect.New(field.Type().Elem()).Elem()
			} else {
				field = field.Elem()
			}
		}
		if field.Kind() == reflect.Struct {
			if err := validateStruct(field.Interface()); err != nil {
				return fmt.Errorf("%s.%s", v.Type().Name(), err)
			}
			continue
		}

		tag := v.Type().Field(i).Tag.Get("required")
		if tag != "true" && tag != "false" {
			return fmt.Errorf("%s's required tag is empty", v.Type().Field(i).Name)
		}
		if tag == "true" && field.IsZero() {
			return fmt.Errorf("required field %s is empty", v.Type().Field(i).Name)
		}
	}

	return nil
}
