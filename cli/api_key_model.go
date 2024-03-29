// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/models"

	"github.com/spf13/cobra"
)

// Schema cli for APIKey

// register flags to command
func registerModelAPIKeyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAPIKeyAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAPIKeyAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelAPIKeyAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAPIKeyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAPIKeyFlags(depth int, m *models.APIKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveAPIKeyAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveAPIKeyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveAPIKeyAttributesFlags(depth int, m *models.APIKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes APIKeyAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.APIKeyAttributes{}
	}

	err, attributesAdded := retrieveModelAPIKeyAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveAPIKeyIDFlags(depth int, m *models.APIKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for APIKeyAttributes

// register flags to command
func registerModelAPIKeyAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAPIKeyAttributesAPIVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyAttributesCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyAttributesLastUsedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyAttributesTokenLastSlice(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyAttributesUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyAttributesUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAPIKeyAttributesAPIVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apiVersionDescription := `The API version associated with this API Key`

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "api_version"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.api_version", cmdPrefix)
	}

	var apiVersionFlagDefault string

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}

func registerAPIKeyAttributesCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdAtDescription := `The time when the API Key was created`

	var createdAtFlagName string
	if cmdPrefix == "" {
		createdAtFlagName = "created_at"
	} else {
		createdAtFlagName = fmt.Sprintf("%v.created_at", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(createdAtFlagName, "", createdAtDescription)

	return nil
}

func registerAPIKeyAttributesLastUsedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastUsedAtDescription := `The last time a request was made to the API using this API Key`

	var lastUsedAtFlagName string
	if cmdPrefix == "" {
		lastUsedAtFlagName = "last_used_at"
	} else {
		lastUsedAtFlagName = fmt.Sprintf("%v.last_used_at", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(lastUsedAtFlagName, "", lastUsedAtDescription)

	return nil
}

func registerAPIKeyAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the API Key`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerAPIKeyAttributesTokenLastSlice(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	tokenLastSliceDescription := `The last 5 characters of the token created for this API Key`

	var tokenLastSliceFlagName string
	if cmdPrefix == "" {
		tokenLastSliceFlagName = "token_last_slice"
	} else {
		tokenLastSliceFlagName = fmt.Sprintf("%v.token_last_slice", cmdPrefix)
	}

	var tokenLastSliceFlagDefault string

	_ = cmd.PersistentFlags().String(tokenLastSliceFlagName, tokenLastSliceFlagDefault, tokenLastSliceDescription)

	return nil
}

func registerAPIKeyAttributesUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updatedAtDescription := `The time when the API Key was updated`

	var updatedAtFlagName string
	if cmdPrefix == "" {
		updatedAtFlagName = "updated_at"
	} else {
		updatedAtFlagName = fmt.Sprintf("%v.updated_at", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(updatedAtFlagName, "", updatedAtDescription)

	return nil
}

func registerAPIKeyAttributesUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var userFlagName string
	if cmdPrefix == "" {
		userFlagName = "user"
	} else {
		userFlagName = fmt.Sprintf("%v.user", cmdPrefix)
	}

	if err := registerModelAPIKeyAttributesUserFlags(depth+1, userFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAPIKeyAttributesFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, apiVersionAdded := retrieveAPIKeyAttributesAPIVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apiVersionAdded

	err, createdAtAdded := retrieveAPIKeyAttributesCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, lastUsedAtAdded := retrieveAPIKeyAttributesLastUsedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastUsedAtAdded

	err, nameAdded := retrieveAPIKeyAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, tokenLastSliceAdded := retrieveAPIKeyAttributesTokenLastSliceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tokenLastSliceAdded

	err, updatedAtAdded := retrieveAPIKeyAttributesUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	err, userAdded := retrieveAPIKeyAttributesUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAdded

	return nil, retAdded
}

func retrieveAPIKeyAttributesAPIVersionFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apiVersionFlagName := fmt.Sprintf("%v.api_version", cmdPrefix)
	if cmd.Flags().Changed(apiVersionFlagName) {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "api_version"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.api_version", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = apiVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIKeyAttributesCreatedAtFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createdAtFlagName := fmt.Sprintf("%v.created_at", cmdPrefix)
	if cmd.Flags().Changed(createdAtFlagName) {

		var createdAtFlagName string
		if cmdPrefix == "" {
			createdAtFlagName = "created_at"
		} else {
			createdAtFlagName = fmt.Sprintf("%v.created_at", cmdPrefix)
		}

		createdAtFlagValueStr, err := cmd.Flags().GetString(createdAtFlagName)
		if err != nil {
			return err, false
		}
		var createdAtFlagValue strfmt.DateTime
		if err := createdAtFlagValue.UnmarshalText([]byte(createdAtFlagValueStr)); err != nil {
			return err, false
		}
		m.CreatedAt = createdAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIKeyAttributesLastUsedAtFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastUsedAtFlagName := fmt.Sprintf("%v.last_used_at", cmdPrefix)
	if cmd.Flags().Changed(lastUsedAtFlagName) {

		var lastUsedAtFlagName string
		if cmdPrefix == "" {
			lastUsedAtFlagName = "last_used_at"
		} else {
			lastUsedAtFlagName = fmt.Sprintf("%v.last_used_at", cmdPrefix)
		}

		lastUsedAtFlagValueStr, err := cmd.Flags().GetString(lastUsedAtFlagName)
		if err != nil {
			return err, false
		}
		var lastUsedAtFlagValue strfmt.DateTime
		if err := lastUsedAtFlagValue.UnmarshalText([]byte(lastUsedAtFlagValueStr)); err != nil {
			return err, false
		}
		m.LastUsedAt = &lastUsedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIKeyAttributesNameFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIKeyAttributesTokenLastSliceFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tokenLastSliceFlagName := fmt.Sprintf("%v.token_last_slice", cmdPrefix)
	if cmd.Flags().Changed(tokenLastSliceFlagName) {

		var tokenLastSliceFlagName string
		if cmdPrefix == "" {
			tokenLastSliceFlagName = "token_last_slice"
		} else {
			tokenLastSliceFlagName = fmt.Sprintf("%v.token_last_slice", cmdPrefix)
		}

		tokenLastSliceFlagValue, err := cmd.Flags().GetString(tokenLastSliceFlagName)
		if err != nil {
			return err, false
		}
		m.TokenLastSlice = tokenLastSliceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIKeyAttributesUpdatedAtFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updatedAtFlagName := fmt.Sprintf("%v.updated_at", cmdPrefix)
	if cmd.Flags().Changed(updatedAtFlagName) {

		var updatedAtFlagName string
		if cmdPrefix == "" {
			updatedAtFlagName = "updated_at"
		} else {
			updatedAtFlagName = fmt.Sprintf("%v.updated_at", cmdPrefix)
		}

		updatedAtFlagValueStr, err := cmd.Flags().GetString(updatedAtFlagName)
		if err != nil {
			return err, false
		}
		var updatedAtFlagValue strfmt.DateTime
		if err := updatedAtFlagValue.UnmarshalText([]byte(updatedAtFlagValueStr)); err != nil {
			return err, false
		}
		m.UpdatedAt = updatedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIKeyAttributesUserFlags(depth int, m *models.APIKeyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userFlagName := fmt.Sprintf("%v.user", cmdPrefix)
	if cmd.Flags().Changed(userFlagName) {
		// info: complex object user APIKeyAttributesUser is retrieved outside this Changed() block
	}
	userFlagValue := m.User
	if swag.IsZero(userFlagValue) {
		userFlagValue = &models.APIKeyAttributesUser{}
	}

	err, userAdded := retrieveModelAPIKeyAttributesUserFlags(depth+1, userFlagValue, userFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAdded
	if userAdded {
		m.User = userFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for APIKeyAttributesUser

// register flags to command
func registerModelAPIKeyAttributesUserFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAPIKeyAttributesUserEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIKeyAttributesUserID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAPIKeyAttributesUserEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := ``

	var emailFlagName string
	if cmdPrefix == "" {
		emailFlagName = "email"
	} else {
		emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
	}

	var emailFlagDefault string

	_ = cmd.PersistentFlags().String(emailFlagName, emailFlagDefault, emailDescription)

	return nil
}

func registerAPIKeyAttributesUserID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAPIKeyAttributesUserFlags(depth int, m *models.APIKeyAttributesUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, emailAdded := retrieveAPIKeyAttributesUserEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, idAdded := retrieveAPIKeyAttributesUserIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveAPIKeyAttributesUserEmailFlags(depth int, m *models.APIKeyAttributesUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	emailFlagName := fmt.Sprintf("%v.email", cmdPrefix)
	if cmd.Flags().Changed(emailFlagName) {

		var emailFlagName string
		if cmdPrefix == "" {
			emailFlagName = "email"
		} else {
			emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
		}

		emailFlagValue, err := cmd.Flags().GetString(emailFlagName)
		if err != nil {
			return err, false
		}
		m.Email = emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIKeyAttributesUserIDFlags(depth int, m *models.APIKeyAttributesUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}
