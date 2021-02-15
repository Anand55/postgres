package main

type DatabaseConfig struct {
	BaseUrl  string `json:"baseUrl"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
	DBName   string `json:"dbname"`
}

//func (d *DatabaseConfig) UnmarshalJSON(e []byte) error {
//	ser := DatabaseConfig{}
//	err := yaml.Unmarshal(e, ser)
//	if err != nil {
//		return err
//	}
//	d.BaseUrl = ser.BaseUrl
//	d.Port = ser.Port
//	d.Password = ser.Password
//	d.User = ser.User
//	d.DBName = ser.DBName
//	return nil
//}
