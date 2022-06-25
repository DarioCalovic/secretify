package config

import (
	"os"

	"github.com/DarioCalovic/secretify"
	flag "github.com/itdistrict/flag"
)

type Configuration struct {
	GlobalPassphraseFile string    `json:"-"`
	SMTP                 *SMTP     `json:"smtp"`
	Meta                 *Meta     `json:"meta"`
	Server               *Server   `json:"server"`
	DB                   *Database `json:"database"`
	Storage              *Storage  `json:"storage"`
	Policy               *Policy   `json:"policy"`
}

type SMTP struct {
	MailJet struct {
		APIKeyPublic  string `json:"api_key_public"`
		APIKeyPrivate string `json:"api_key_private"`
	} `json:"mailjet"`
}

type Meta struct {
	// UIURL  string `json:"ui_url"`
	Hoster struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"hoster"`
}

type Policy struct {
	Identifier struct {
		Size int `json:"size"`
	} `json:"identifier"`
	Mail struct {
		Enabled bool `json:"enabled"`
	} `json:"mail"`
	Webhook struct {
		Enabled bool `json:"enabled"`
	} `json:"webhook"`
	Secret struct {
		MaxLength int `json:"max_length"`
	} `json:"secret"`
	Passphrase struct {
		Required bool `json:"required"`
	} `json:"passphrase"`
	Storage struct {
		FileSystem struct {
			MaxFileSize           uint   `json:"max_filesize"`
			AllowedFileExtensions string `json:"allowed_file_extensions"`
		} `json:"filesystem"`
	} `json:"storage"`
}

type Server struct {
	Address      string `json:"address"`
	BasePath     string `json:"base_path"`
	Debug        bool   `json:"debug"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	Auth         *Auth
}

type Auth struct {
	Enabled bool
	JWT     *JWT
}

// JWT holds data necessary for JWT configuration
type JWT struct {
	MinSecretLength  int    `yaml:"min_secret_length,omitempty"`
	DurationMinutes  int    `yaml:"duration_minutes,omitempty"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `yaml:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

type Database struct {
	ConnectionURL string `json:"connection_url"`
}

type Storage struct {
	Provider   string     `json:"provider"`
	FileSystem FileSystem `json:"filesystem"`
}

type FileSystem struct {
	Location string `json:"location"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		SMTP: &SMTP{},
		Meta: &Meta{},
		Server: &Server{
			Auth: &Auth{},
		},
		DB:      &Database{},
		Storage: &Storage{},
		Policy:  &Policy{},
	}
}

func Load() (*Configuration, error) {
	fs := new(flag.FlagSet)
	cfg := ParseFlags(fs)

	err := cfg.Validate()

	return cfg, err
}

func (c *Configuration) Validate() error {
	if c.Server.BasePath == "/" {
		c.Server.BasePath = ""
	}
	return nil
}

func ParseFlags(fs *flag.FlagSet) *Configuration {
	cfg := NewConfiguration()

	// Database
	fs.StringVar(&cfg.DB.ConnectionURL, "Database_ConnectionURL", "secretify.db", "The sqlite database connection url string (path to database file)")

	// Storage
	fs.StringVar(&cfg.Storage.Provider, "Storage_Provider", "filesystem", "Storage provider name")
	fs.StringVar(&cfg.Storage.FileSystem.Location, "Storage_Filesystem_Location", "./_files", "Storage filesystem location where the files will be stored")

	// Server
	fs.StringVar(&cfg.Server.Address, "Server_Address", ":8800", "The http port the api will listen to")
	fs.StringVar(&cfg.Server.BasePath, "Server_BasePath", "/", "Base path")
	fs.BoolVar(&cfg.Server.Debug, "Server_Debug", false, "Enable api debug mode")
	fs.IntVar(&cfg.Server.ReadTimeout, "Server_ReadTimeout", 10, "Read timeout in seconds for http requests")
	fs.IntVar(&cfg.Server.WriteTimeout, "Server_WriteTimeout", 100, "Write timeout in seconds for http requests")

	// Policy
	fs.IntVar(&cfg.Policy.Identifier.Size, "Policy_Identifier_Size", 18, "Size of the identifier")
	fs.BoolVar(&cfg.Policy.Mail.Enabled, "Policy_Mail_Enabled", true, "Enable sending emails or not")
	fs.BoolVar(&cfg.Policy.Webhook.Enabled, "Policy_Webhook_Enabled", true, "Enable issueing webhooks")
	fs.IntVar(&cfg.Policy.Secret.MaxLength, "Policy_Secret_MaxLength", 500, "Secret's max length")
	fs.BoolVar(&cfg.Policy.Passphrase.Required, "Policy_Passphrase_Required", false, "Passphrase required or not")
	fs.UintVar(&cfg.Policy.Storage.FileSystem.MaxFileSize, "Policy_Storage_Filesystem_MaxFileSize", secretify.MaxFileSize, "Storage max size per file")
	fs.StringVar(&cfg.Policy.Storage.FileSystem.AllowedFileExtensions, "Policy_Storage_Filesystem_AllowedFileExtensions", secretify.AllowedFileExtensions, "Allowed file extensions (comma separated), leave empty if all types should be allowed")

	// Meta
	fs.StringVar(&cfg.Meta.Hoster.Name, "Meta_Hoster_Name", "", "Meta information about the hoster name")
	fs.StringVar(&cfg.Meta.Hoster.Address, "Meta_Hoster_Address", "", "Meta information about the hoster address")

	// SMTP
	fs.StringVar(&cfg.SMTP.MailJet.APIKeyPublic, "SMTP_MailJet_APIKey_Public", "c4bf39bd29338413af37451bbd2ac507", "MailJet public api key")
	fs.StringVar(&cfg.SMTP.MailJet.APIKeyPrivate, "SMTP_MailJet_APIKey_Private", "ebf4d1945f419fb055b003b157b126a0", "MailJet private api key")

	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	return cfg
}
