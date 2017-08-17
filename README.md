# json-conf
A library that reads app's json-formatted config files and wipes out comments

# Features
1. Support comments like "//" at the beginning or end of lines.
2. Support comments like "/* ... */" in multiple lines.

# Non-features
1. Don't support Embeded comments like "/* ... /* ... */... */" in single or multiple lines.

# Example
Json-formmated config file is:

	{
		// all field.
		"timeout":100, // time out in http
		/*"ip":"127.0.0.1",
		"port":8000,*/
		"redis":"127.0.0.1:6379",
		"mysql":"127.0.0.1:3306"
	}
Your AppConfig struct is defined as:

	type MyConf struct {
		Timeout int `json:"timeout"`
		Ip string `json:"ip"`
		Port int `json:"port"`
		Redis string `json:"redis"`
		Mysql string `json:"mysql"`
		Log string `json:"log"`
	}
After jsonconf.Unmarshal(), print(myConf) gets this:

{100  0 127.0.0.1:6379 127.0.0.1:3306 }
