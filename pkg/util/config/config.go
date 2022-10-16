package config

import (
	"errors"
	"os"
	"strings"

	"github.com/DarioCalovic/secretify"
	flag "github.com/itdistrict/flag"
)

type Configuration struct {
	GlobalPassphraseFile string    `json:"-"`
	Meta                 *Meta     `json:"meta"`
	Server               *Server   `json:"server"`
	DB                   *Database `json:"database"`
	Storage              *Storage  `json:"storage"`
	Policy               *Policy   `json:"policy"`
	Outlook              *Outlook  `json:"outlook"`
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
	Webhook struct {
		Enabled bool `json:"enabled"`
	} `json:"webhook"`
	Secret struct {
		MaxLength      int    `json:"max_length"`
		RevealDuration string `json:"reveal_duration"`
	} `json:"secret"`
	Passphrase struct {
		Required bool `json:"required"`
	} `json:"passphrase"`
	Storage struct {
		Enabled    bool `json:"enabled"`
		FileSystem struct {
			MaxFileSize           uint   `json:"max_filesize"`
			AllowedFileExtensions string `json:"allowed_file_extensions"`
		} `json:"filesystem"`
	} `json:"storage"`
	Security Security `json:"security"`
}

type Security struct {
	CORS struct {
		allowedOrigins string
		AllowedOrigins []string `json:"allowed_origins"`
	} `json:"cors"`
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
	Enabled    bool       `json:"enabled"`
	Provider   string     `json:"provider"`
	FileSystem FileSystem `json:"filesystem"`
}

type FileSystem struct {
	Location string `json:"location"`
}

type Outlook struct {
	Enabled bool   `json:"enabled"`
	AppID   string `json:"app_id"`
	UIURL   string `json:"ui_url"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Meta: &Meta{},
		Server: &Server{
			Auth: &Auth{},
		},
		DB:      &Database{},
		Storage: &Storage{},
		Policy:  &Policy{},
		Outlook: &Outlook{},
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
	// Outlook
	if c.Outlook.Enabled && c.Outlook.AppID == "" {
		return errors.New("no outlook app id provided")
	}
	return nil
}

func ParseFlags(fs *flag.FlagSet) *Configuration {
	cfg := NewConfiguration()

	// Database
	fs.StringVar(&cfg.DB.ConnectionURL, "Database_ConnectionURL", "./db/secretify.db", "The sqlite database connection url string (path to database file)")

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
	fs.IntVar(&cfg.Policy.Secret.MaxLength, "Policy_Secret_MaxLength", 500, "Secret's max length")
	fs.StringVar(&cfg.Policy.Secret.RevealDuration, "Policy_Secret_RevealDuration", "60s", "Amount of time the secret will be shown when revealing on the UI")
	fs.BoolVar(&cfg.Policy.Passphrase.Required, "Policy_Passphrase_Required", false, "Passphrase required or not")
	fs.BoolVar(&cfg.Policy.Storage.Enabled, "Policy_Storage_Enabled", false, "Enable sharing files")
	fs.UintVar(&cfg.Policy.Storage.FileSystem.MaxFileSize, "Policy_Storage_Filesystem_MaxFileSize", secretify.MaxFileSize, "Storage max size per file")
	fs.StringVar(&cfg.Policy.Storage.FileSystem.AllowedFileExtensions, "Policy_Storage_Filesystem_AllowedFileExtensions", secretify.AllowedFileExtensions, "Allowed file extensions (comma separated), leave empty if all types should be allowed")

	// Policy Security
	fs.StringVar(&cfg.Policy.Security.CORS.allowedOrigins, "Policy_Security_CORS_AllowedOrigins", "*", "Set CORS AllowedOrigins settings for security reasons (comma separated)")

	// Outlook
	fs.BoolVar(&cfg.Outlook.Enabled, "Outlook_Enabled", false, "Enable outlook integration")
	fs.StringVar(&cfg.Outlook.AppID, "Outlook_AppID", "", "Outlook add-in app id")
	fs.StringVar(&cfg.Outlook.UIURL, "Outlook_UI_URL", "https://localhost:3000", "Outlook add-in ui url")

	// Meta
	fs.StringVar(&cfg.Meta.Hoster.Name, "Meta_Hoster_Name", "", "Meta information about the hoster name")
	fs.StringVar(&cfg.Meta.Hoster.Address, "Meta_Hoster_Address", "", "Meta information about the hoster address")

	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	// Transform some parameters
	cfg.Policy.Security.CORS.AllowedOrigins = strings.Split(cfg.Policy.Security.CORS.allowedOrigins, ",")

	return cfg
}
