package msq

func initDBTable() (err error) {
	if err = createUserTable(); err != nil {
		return err
	}
	if err = createCommunityTable(); err != nil {
		return err
	}
	return nil
}

func createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS user(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    uuid BIGINT UNIQUE,
    username VARCHAR(16) NOT NULL UNIQUE ,
    hash VARCHAR(64) NOT NULL ,
    email VARCHAR(64),
    INDEX idx_uuid (uuid),
    INDEX idx_username (username)
)ENGINE = InnoDB DEFAULT CHARSET utf8mb4`
	_, err := db.Exec(query)
	return err
}

func createCommunityTable() error {
	query := `CREATE TABLE IF NOT EXISTS community(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(16) NOT NULL UNIQUE ,
    admin_uuid BIGINT,
    administrator VARCHAR(16),
    introduction VARCHAR(512) NOT NULL,
    INDEX idx_name (name)
)ENGINE=InnoDB DEFAULT CHARSET utf8mb4`
	_, err := db.Exec(query)
	return err
}
