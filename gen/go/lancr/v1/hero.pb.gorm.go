package v1

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	gorm "gorm.io/gorm"
	strings "strings"
)

type ExperienceORM struct {
	Company   string
	EndDate   string
	HeroId    *string
	Id        string
	Position  string
	StartDate string
}

// TableName overrides the default tablename generated by GORM
func (ExperienceORM) TableName() string {
	return "experiences"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Experience) ToORM(ctx context.Context) (ExperienceORM, error) {
	to := ExperienceORM{}
	var err error
	if prehook, ok := interface{}(m).(ExperienceWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Position = m.Position
	to.Company = m.Company
	to.StartDate = m.StartDate
	to.EndDate = m.EndDate
	if posthook, ok := interface{}(m).(ExperienceWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ExperienceORM) ToPB(ctx context.Context) (Experience, error) {
	to := Experience{}
	var err error
	if prehook, ok := interface{}(m).(ExperienceWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Position = m.Position
	to.Company = m.Company
	to.StartDate = m.StartDate
	to.EndDate = m.EndDate
	if posthook, ok := interface{}(m).(ExperienceWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Experience the arg will be the target, the caller the one being converted from

// ExperienceBeforeToORM called before default ToORM code
type ExperienceWithBeforeToORM interface {
	BeforeToORM(context.Context, *ExperienceORM) error
}

// ExperienceAfterToORM called after default ToORM code
type ExperienceWithAfterToORM interface {
	AfterToORM(context.Context, *ExperienceORM) error
}

// ExperienceBeforeToPB called before default ToPB code
type ExperienceWithBeforeToPB interface {
	BeforeToPB(context.Context, *Experience) error
}

// ExperienceAfterToPB called after default ToPB code
type ExperienceWithAfterToPB interface {
	AfterToPB(context.Context, *Experience) error
}

type HeroORM struct {
	Bio            string
	Email          string
	Experience     []*ExperienceORM `gorm:"foreignKey:HeroId;references:Id"`
	FirebaseId     string
	Id             string
	Language       string
	Level          uint32
	Linkedin       string
	PhoneNumber    string
	ProfilePicture string
	Rating         float32
	Region         string
	SubWorkType    string
	TeamId         *string
	Twitter        string
	UserName       string
	Visibility     string
	Website        string
	WorkType       string
	Xp             string
}

// TableName overrides the default tablename generated by GORM
func (HeroORM) TableName() string {
	return "heros"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Hero) ToORM(ctx context.Context) (HeroORM, error) {
	to := HeroORM{}
	var err error
	if prehook, ok := interface{}(m).(HeroWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.FirebaseId = m.FirebaseId
	to.Visibility = m.Visibility
	to.ProfilePicture = m.ProfilePicture
	to.Email = m.Email
	to.PhoneNumber = m.PhoneNumber
	to.UserName = m.UserName
	to.Bio = m.Bio
	to.WorkType = m.WorkType
	to.SubWorkType = m.SubWorkType
	to.Twitter = m.Twitter
	to.Linkedin = m.Linkedin
	to.Website = m.Website
	for _, v := range m.Experience {
		if v != nil {
			if tempExperience, cErr := v.ToORM(ctx); cErr == nil {
				to.Experience = append(to.Experience, &tempExperience)
			} else {
				return to, cErr
			}
		} else {
			to.Experience = append(to.Experience, nil)
		}
	}
	to.Rating = m.Rating
	to.Xp = m.Xp
	to.Region = m.Region
	to.Language = m.Language
	to.Level = m.Level
	if posthook, ok := interface{}(m).(HeroWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *HeroORM) ToPB(ctx context.Context) (Hero, error) {
	to := Hero{}
	var err error
	if prehook, ok := interface{}(m).(HeroWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.FirebaseId = m.FirebaseId
	to.Visibility = m.Visibility
	to.ProfilePicture = m.ProfilePicture
	to.Email = m.Email
	to.PhoneNumber = m.PhoneNumber
	to.UserName = m.UserName
	to.Bio = m.Bio
	to.WorkType = m.WorkType
	to.SubWorkType = m.SubWorkType
	to.Twitter = m.Twitter
	to.Linkedin = m.Linkedin
	to.Website = m.Website
	for _, v := range m.Experience {
		if v != nil {
			if tempExperience, cErr := v.ToPB(ctx); cErr == nil {
				to.Experience = append(to.Experience, &tempExperience)
			} else {
				return to, cErr
			}
		} else {
			to.Experience = append(to.Experience, nil)
		}
	}
	to.Rating = m.Rating
	to.Xp = m.Xp
	to.Region = m.Region
	to.Language = m.Language
	to.Level = m.Level
	if posthook, ok := interface{}(m).(HeroWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Hero the arg will be the target, the caller the one being converted from

// HeroBeforeToORM called before default ToORM code
type HeroWithBeforeToORM interface {
	BeforeToORM(context.Context, *HeroORM) error
}

// HeroAfterToORM called after default ToORM code
type HeroWithAfterToORM interface {
	AfterToORM(context.Context, *HeroORM) error
}

// HeroBeforeToPB called before default ToPB code
type HeroWithBeforeToPB interface {
	BeforeToPB(context.Context, *Hero) error
}

// HeroAfterToPB called after default ToPB code
type HeroWithAfterToPB interface {
	AfterToPB(context.Context, *Hero) error
}

// DefaultCreateExperience executes a basic gorm create call
func DefaultCreateExperience(ctx context.Context, in *Experience, db *gorm.DB) (*Experience, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ExperienceORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadExperience(ctx context.Context, in *Experience, db *gorm.DB) (*Experience, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ExperienceORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ExperienceORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ExperienceORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteExperience(ctx context.Context, in *Experience, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ExperienceORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ExperienceORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteExperienceSet(ctx context.Context, in []*Experience, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&ExperienceORM{})).(ExperienceORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ExperienceORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ExperienceORM{})).(ExperienceORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ExperienceORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Experience, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Experience, *gorm.DB) error
}

// DefaultStrictUpdateExperience clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateExperience(ctx context.Context, in *Experience, db *gorm.DB) (*Experience, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateExperience")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ExperienceORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type ExperienceORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchExperience executes a basic gorm update call with patch behavior
func DefaultPatchExperience(ctx context.Context, in *Experience, updateMask *field_mask.FieldMask, db *gorm.DB) (*Experience, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Experience
	var err error
	if hook, ok := interface{}(&pbObj).(ExperienceWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadExperience(ctx, &Experience{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ExperienceWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskExperience(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ExperienceWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateExperience(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ExperienceWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ExperienceWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Experience, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ExperienceWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Experience, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ExperienceWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Experience, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ExperienceWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Experience, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetExperience executes a bulk gorm update call with patch behavior
func DefaultPatchSetExperience(ctx context.Context, objects []*Experience, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Experience, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Experience, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchExperience(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskExperience patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskExperience(ctx context.Context, patchee *Experience, patcher *Experience, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Experience, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Position" {
			patchee.Position = patcher.Position
			continue
		}
		if f == prefix+"Company" {
			patchee.Company = patcher.Company
			continue
		}
		if f == prefix+"StartDate" {
			patchee.StartDate = patcher.StartDate
			continue
		}
		if f == prefix+"EndDate" {
			patchee.EndDate = patcher.EndDate
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListExperience executes a gorm list call
func DefaultListExperience(ctx context.Context, db *gorm.DB) ([]*Experience, error) {
	in := Experience{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ExperienceORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExperienceORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Experience{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ExperienceORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ExperienceORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]ExperienceORM) error
}

// DefaultCreateHero executes a basic gorm create call
func DefaultCreateHero(ctx context.Context, in *Hero, db *gorm.DB) (*Hero, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type HeroORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadHero(ctx context.Context, in *Hero, db *gorm.DB) (*Hero, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := HeroORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(HeroORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type HeroORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteHero(ctx context.Context, in *Hero, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&HeroORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type HeroORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteHeroSet(ctx context.Context, in []*Hero, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&HeroORM{})).(HeroORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&HeroORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&HeroORM{})).(HeroORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type HeroORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Hero, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Hero, *gorm.DB) error
}

// DefaultStrictUpdateHero clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateHero(ctx context.Context, in *Hero, db *gorm.DB) (*Hero, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateHero")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &HeroORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	filterExperience := ExperienceORM{}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	filterExperience.HeroId = new(string)
	*filterExperience.HeroId = ormObj.Id
	if err = db.Where(filterExperience).Delete(ExperienceORM{}).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type HeroORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchHero executes a basic gorm update call with patch behavior
func DefaultPatchHero(ctx context.Context, in *Hero, updateMask *field_mask.FieldMask, db *gorm.DB) (*Hero, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Hero
	var err error
	if hook, ok := interface{}(&pbObj).(HeroWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadHero(ctx, &Hero{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(HeroWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskHero(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(HeroWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateHero(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(HeroWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type HeroWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Hero, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type HeroWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Hero, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type HeroWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Hero, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type HeroWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Hero, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetHero executes a bulk gorm update call with patch behavior
func DefaultPatchSetHero(ctx context.Context, objects []*Hero, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Hero, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Hero, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchHero(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskHero patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskHero(ctx context.Context, patchee *Hero, patcher *Hero, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Hero, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedName bool
	var updatedLocation bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"FirebaseId" {
			patchee.FirebaseId = patcher.FirebaseId
			continue
		}
		if !updatedName && strings.HasPrefix(f, prefix+"Name.") {
			if patcher.Name == nil {
				patchee.Name = nil
				continue
			}
			if patchee.Name == nil {
				patchee.Name = &Name{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Name."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.Name, patchee.Name, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Name" {
			updatedName = true
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Visibility" {
			patchee.Visibility = patcher.Visibility
			continue
		}
		if f == prefix+"ProfilePicture" {
			patchee.ProfilePicture = patcher.ProfilePicture
			continue
		}
		if f == prefix+"Email" {
			patchee.Email = patcher.Email
			continue
		}
		if f == prefix+"PhoneNumber" {
			patchee.PhoneNumber = patcher.PhoneNumber
			continue
		}
		if !updatedLocation && strings.HasPrefix(f, prefix+"Location.") {
			if patcher.Location == nil {
				patchee.Location = nil
				continue
			}
			if patchee.Location == nil {
				patchee.Location = &Location{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Location."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.Location, patchee.Location, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Location" {
			updatedLocation = true
			patchee.Location = patcher.Location
			continue
		}
		if f == prefix+"UserName" {
			patchee.UserName = patcher.UserName
			continue
		}
		if f == prefix+"Bio" {
			patchee.Bio = patcher.Bio
			continue
		}
		if f == prefix+"WorkType" {
			patchee.WorkType = patcher.WorkType
			continue
		}
		if f == prefix+"SubWorkType" {
			patchee.SubWorkType = patcher.SubWorkType
			continue
		}
		if f == prefix+"Twitter" {
			patchee.Twitter = patcher.Twitter
			continue
		}
		if f == prefix+"Linkedin" {
			patchee.Linkedin = patcher.Linkedin
			continue
		}
		if f == prefix+"Website" {
			patchee.Website = patcher.Website
			continue
		}
		if f == prefix+"Experience" {
			patchee.Experience = patcher.Experience
			continue
		}
		if f == prefix+"Rating" {
			patchee.Rating = patcher.Rating
			continue
		}
		if f == prefix+"Xp" {
			patchee.Xp = patcher.Xp
			continue
		}
		if f == prefix+"Region" {
			patchee.Region = patcher.Region
			continue
		}
		if f == prefix+"Language" {
			patchee.Language = patcher.Language
			continue
		}
		if f == prefix+"Level" {
			patchee.Level = patcher.Level
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListHero executes a gorm list call
func DefaultListHero(ctx context.Context, db *gorm.DB) ([]*Hero, error) {
	in := Hero{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []HeroORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(HeroORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Hero{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type HeroORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type HeroORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]HeroORM) error
}
