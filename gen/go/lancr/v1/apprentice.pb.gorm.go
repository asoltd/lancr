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

type ApprenticeORM struct {
	FavoriteTo string
	Id         string
	Mentor     string
	Rate       float32
}

// TableName overrides the default tablename generated by GORM
func (ApprenticeORM) TableName() string {
	return "apprentices"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Apprentice) ToORM(ctx context.Context) (ApprenticeORM, error) {
	to := ApprenticeORM{}
	var err error
	if prehook, ok := interface{}(m).(ApprenticeWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Rate = m.Rate
	to.Mentor = m.Mentor
	to.FavoriteTo = m.FavoriteTo
	// Repeated type string is not an ORMable message type
	if posthook, ok := interface{}(m).(ApprenticeWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ApprenticeORM) ToPB(ctx context.Context) (Apprentice, error) {
	to := Apprentice{}
	var err error
	if prehook, ok := interface{}(m).(ApprenticeWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Rate = m.Rate
	to.Mentor = m.Mentor
	to.FavoriteTo = m.FavoriteTo
	// Repeated type string is not an ORMable message type
	if posthook, ok := interface{}(m).(ApprenticeWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Apprentice the arg will be the target, the caller the one being converted from

// ApprenticeBeforeToORM called before default ToORM code
type ApprenticeWithBeforeToORM interface {
	BeforeToORM(context.Context, *ApprenticeORM) error
}

// ApprenticeAfterToORM called after default ToORM code
type ApprenticeWithAfterToORM interface {
	AfterToORM(context.Context, *ApprenticeORM) error
}

// ApprenticeBeforeToPB called before default ToPB code
type ApprenticeWithBeforeToPB interface {
	BeforeToPB(context.Context, *Apprentice) error
}

// ApprenticeAfterToPB called after default ToPB code
type ApprenticeWithAfterToPB interface {
	AfterToPB(context.Context, *Apprentice) error
}

// DefaultCreateApprentice executes a basic gorm create call
func DefaultCreateApprentice(ctx context.Context, in *Apprentice, db *gorm.DB) (*Apprentice, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ApprenticeORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadApprentice(ctx context.Context, in *Apprentice, db *gorm.DB) (*Apprentice, error) {
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
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ApprenticeORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ApprenticeORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ApprenticeORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteApprentice(ctx context.Context, in *Apprentice, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ApprenticeORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ApprenticeORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteApprenticeSet(ctx context.Context, in []*Apprentice, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&ApprenticeORM{})).(ApprenticeORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ApprenticeORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ApprenticeORM{})).(ApprenticeORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ApprenticeORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Apprentice, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Apprentice, *gorm.DB) error
}

// DefaultStrictUpdateApprentice clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateApprentice(ctx context.Context, in *Apprentice, db *gorm.DB) (*Apprentice, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateApprentice")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ApprenticeORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithAfterStrictUpdateSave); ok {
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

type ApprenticeORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchApprentice executes a basic gorm update call with patch behavior
func DefaultPatchApprentice(ctx context.Context, in *Apprentice, updateMask *field_mask.FieldMask, db *gorm.DB) (*Apprentice, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Apprentice
	var err error
	if hook, ok := interface{}(&pbObj).(ApprenticeWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadApprentice(ctx, &Apprentice{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ApprenticeWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskApprentice(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ApprenticeWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateApprentice(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ApprenticeWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ApprenticeWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Apprentice, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Apprentice, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Apprentice, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Apprentice, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetApprentice executes a bulk gorm update call with patch behavior
func DefaultPatchSetApprentice(ctx context.Context, objects []*Apprentice, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Apprentice, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Apprentice, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchApprentice(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskApprentice patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskApprentice(ctx context.Context, patchee *Apprentice, patcher *Apprentice, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Apprentice, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedWorkingHours bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Rate" {
			patchee.Rate = patcher.Rate
			continue
		}
		if !updatedWorkingHours && strings.HasPrefix(f, prefix+"WorkingHours.") {
			if patcher.WorkingHours == nil {
				patchee.WorkingHours = nil
				continue
			}
			if patchee.WorkingHours == nil {
				patchee.WorkingHours = &WorkingHours{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"WorkingHours."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.WorkingHours, patchee.WorkingHours, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"WorkingHours" {
			updatedWorkingHours = true
			patchee.WorkingHours = patcher.WorkingHours
			continue
		}
		if f == prefix+"Mentor" {
			patchee.Mentor = patcher.Mentor
			continue
		}
		if f == prefix+"FavoriteTo" {
			patchee.FavoriteTo = patcher.FavoriteTo
			continue
		}
		if f == prefix+"WorkingDays" {
			patchee.WorkingDays = patcher.WorkingDays
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListApprentice executes a gorm list call
func DefaultListApprentice(ctx context.Context, db *gorm.DB) ([]*Apprentice, error) {
	in := Apprentice{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ApprenticeORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ApprenticeORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Apprentice{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ApprenticeORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ApprenticeORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]ApprenticeORM) error
}
