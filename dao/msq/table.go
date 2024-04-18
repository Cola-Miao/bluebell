package msq

func initDBTable() (err error) {
	if err = createUserTable(); err != nil {
		return err
	}
	if err = createCommunityTable(); err != nil {
		return err
	}
	if err = createArticleTable(); err != nil {
		return err
	}
	return nil
}

func createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS user(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    uuid BIGINT UNIQUE,
    username VARCHAR(32) NOT NULL UNIQUE ,
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
    name VARCHAR(32) NOT NULL UNIQUE ,
    admin_uuid BIGINT,
    administrator VARCHAR(32),
    introduction VARCHAR(1024) NOT NULL,
    INDEX idx_name (name)
)ENGINE=InnoDB DEFAULT CHARSET utf8mb4`
	_, err := db.Exec(query)
	return err
}

func createArticleTable() error {
	query := `CREATE TABLE IF NOT EXISTS article(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
	uuid BIGINT NOT NULL UNIQUE,
	community_id BIGINT NOT NULL ,
	author_uuid BIGINT NOT NULL ,
	author VARCHAR(32) NOT NULL ,
	title VARCHAR(128) NOT NULL ,
	content TEXT NOT NULL ,
	introduction VARCHAR(512) NOT NULL ,
	create_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	update_at DATETIME DEFAULT CURRENT_TIMESTAMP 
ON UPDATE CURRENT_TIMESTAMP,
	INDEX idx_uuid(uuid),
	INDEX idx_community_id(community_id)
)ENGINE =InnoDB DEFAULT CHARSET utf8mb4`
	_, err := db.Exec(query)
	return err
}
