package bot

import "time"

type (

	// Account implements trading account functionality. It may
	// represent a single account, several accounts, a 'paper'
	// trading account, or any combination.
	//
	// This allows for simple replacement of accounts for testing,
	// scaling of bot activity, accessing accounts with different
	// currencies, etc.
	Account interface{}

	// BotConfig implements bot configuration functionality,
	// such as settings, limits, and other settings.
	//
	// Authentication, Permissions, and Account information are
	// not included or accessible from BotConfig.
	BotConfig interface{}

	// Permissions maintains a mapping of allowed scopes
	// for the current user.
	Permissions interface{}

	// Bot implements the bot. It can take input as rules and conditions.
	// It can be scheduled, managed, and switched on or off.
	Bot interface{}

	// Setting is an individual BotConfig object
	Setting interface{}

	// token is the secure token associated with the current user.
	token struct{}

	// permissionMap is a map of permissions
	permissionMap map[int]bool
)

type bot struct {
	// LinkedAccount represents account associated with this bot.
	// It may actually represent a set of accounts where the
	// Trades are duplicated according to percentage rules
	// within that account.
	LinkedAccount Account
	Configuration config
	Perm          permissions

	// rs ruleSet
}

type permissions struct {
	LoginName  string
	LoginToken token
	PMap       permissionMap
}

type config struct {
	AllowPTD       bool
	TradeStartTime time.Time
	TradeEndTime   time.Time
	MaxPercentage  float64
	MinSize        float64
	MaxSize        float64
}

func (c *config) Set(name string, s Setting) error {
	return nil
}

func (c *config) Get(name string) (Setting, error) {
	return nil, nil
}

func (c *config) Load(name string) (BotConfig, error) {
	return nil, nil
}

func (c *config) Save(name string) error {
	return nil
}
