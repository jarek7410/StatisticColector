package dbStats

//TODO: DELETE
//
//
//
//import (
//	"gorm.io/gorm"
//	"log"
//)
//
//type Repo struct {
//	db *gorm.DB
//	//ctx context.Context
//}
//
//func NewSQLGORMRepository(db *gorm.DB) *Repo {
//	return &Repo{
//		db: db,
//	}
//}
//func (r *Repo) Migrate() error {
//	log.Println("start migrate")
//	return r.db.AutoMigrate(&User{}, &Stat{})
//}
//
//func (r *Repo) CrateUser(user User) (*User, error) {
//	result := r.db.Create(&user)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return &user, nil
//}
//func (r *Repo) CrateStat(user User, stat Stat) (*Stat, error) {
//	result := r.db.First(&user, user.ID)
//	if result.Error != nil {
//		log.Println("db user not found")
//		return nil, errNoUser
//
//	}
//	stat.UserID = user.ID
//	result = r.db.Create(&stat)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return &stat, nil
//}
//
//func (r *Repo) GetUsers() ([]User, error) {
//	var users []User
//	if err := r.db.Find(&users).Error; err != nil {
//		return nil, err
//	}
//	return users, nil
//}
//func (r *Repo) GerUser(userId uint) (*User, error) {
//	var user User
//	if err := r.db.First(&user, userId).Error; err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
//func (r *Repo) GetStats(user User) ([]Stat, error) {
//	var retrieveUser User
//	if err := r.db.Preload("Stats").
//		First(&retrieveUser, user.ID).Error; err != nil {
//		return nil, err
//	}
//	return retrieveUser.Stats, nil
//}
//func (r *Repo) DeleteStat(stat Stat) error {
//	if err := r.db.
//		Delete(&stat).Error; err != nil {
//		return err
//	}
//	return nil
//}
//func (r *Repo) KillUser(user User) error {
//	log.Println("illegal")
//	return errNoUser
//}
