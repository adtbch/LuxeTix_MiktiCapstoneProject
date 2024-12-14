package entity

type User struct {
	ID        int  `json:"id"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Gender    string `json:"gender"`
	CreatedAt int  `json:"created_at"`
	UpdatedAt int  `json:"updated_at"`
}
// CREATE TABLE IF NOT EXISTS users (
//     ID INT AUTO_INCREMENT PRIMARY KEY,
//     Fullname VARCHAR(255) NOT NULL,
//     Username VARCHAR(255) NOT NULL UNIQUE,
//     Email VARCHAR(255) NOT NULL UNIQUE,
//     Password VARCHAR(255) NOT NULL,
//     Role VARCHAR(255) NOT NULL,
//     Gender VARCHAR(255) NOT NULL,
//     Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     Updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );