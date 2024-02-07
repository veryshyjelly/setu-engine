package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"setu-engine/bridge"
	"setu-engine/models"
	"time"
)

type Service interface {
	CreateBridge(bridge models.Bridge) error
	DeleteBridge(fromChatID string, toChatID string) error
	GetBridge(fromChatID string) ([]models.Bridge, error)
	GetAllBridges() ([]models.Bridge, error)
}

type service struct {
	bridgeService bridge.Service
	db            *gorm.DB
}

func Connect(sr bridge.Service) Service {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open("maya_setu.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Bridge{})
	if err != nil {
		panic(err)
	}

	ser := &service{db: db, bridgeService: sr}

	br, err := ser.GetAllBridges()
	if err != nil {
		panic(err)
	}

	for _, b := range br {
		sr.Subscribe() <- b
	}

	return ser
}

func (s *service) CreateBridge(bridge models.Bridge) error {
	// Check if the bridge already exists
	br, err := s.GetBridge(bridge.FromChatID)
	if err != nil {
		return err
	}
	for _, b := range br {
		if b.SecondChatID == bridge.SecondChatID {
			return errors.New("bridge already exists")
		}
	}
	// Check if the bridge already exists
	br, err = s.GetBridge(bridge.SecondChatID)
	if err != nil {
		return err
	}
	for _, b := range br {
		if b.SecondChatID == bridge.FromChatID {
			return errors.New("bridge already exists")
		}
	}
	log.Println("CREATING BRIDGE BETWEEN:", bridge.FromChatID, "and", bridge.SecondChatID)
	err = s.db.Create(&bridge).Error
	if err != nil {
		log.Println("ERROR CREATING BRIDGE:", err)
		return err
	}
	s.bridgeService.Subscribe() <- bridge
	return nil
}

func (s *service) DeleteBridge(firstChatId string, secondChatId string) error {
	return s.db.Where("first_chat_id = ? AND second_chat_id = ?", firstChatId, secondChatId).Delete(&models.Bridge{}).Error
}

func (s *service) GetBridge(fromChatID string) ([]models.Bridge, error) {
	var bridges []models.Bridge
	s.db.Where("from_chat_id = ? or second_chat_id = ?", fromChatID, fromChatID).Find(&bridges)
	return bridges, nil
}

func (s *service) GetAllBridges() ([]models.Bridge, error) {
	var bridges []models.Bridge
	s.db.Find(&bridges)
	return bridges, nil
}