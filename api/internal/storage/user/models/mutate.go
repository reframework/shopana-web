package userStorageModel

type MutateModel struct {
	Email       string  `db:"email" goqu:"omitempty"`
	FirstName   string  `db:"first_name" goqu:"omitempty"`
	IsBlocked   *bool   `db:"is_blocked" goqu:"omitnil"`
	IsVerified  *bool   `db:"is_verified" goqu:"omitnil"`
	Language    string  `db:"language" goqu:"omitempty"`
	LastName    string  `db:"last_name" goqu:"omitempty"`
	Password    string  `db:"password" goqu:"omitempty"`
	PhoneNumber *string `db:"phone_number" goqu:"omitnil"`
	Timezone    string  `db:"timezone" goqu:"omitempty"`
}
