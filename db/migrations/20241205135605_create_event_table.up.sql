BEGIN;

CREATE TABLE IF NOT EXISTS events (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,
    Title VARCHAR(255) NOT NULL UNIQUE,
    Description VARCHAR(255) NOT NULL,
    Date DATE NOT NULL,
    Price INT NOT NULL,
    Quantity INT NOT NULL,
    Time TIME NOT NULL,
    Location VARCHAR(255) NOT NULL,
    StatusEvent ENUM('available', 'unavailable') NOT NULL,
    StatusRequest ENUM('pending', 'unpaid', 'accepted', 'rejected') NOT NULL,
    Category VARCHAR(255) NOT NULL,
    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

COMMIT;
