// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/models"

	"github.com/spf13/cobra"
)

// Schema cli for UserTeam

// register flags to command
func registerModelUserTeamFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserTeamAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserTeamAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelUserTeamAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserTeamID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelUserTeamFlags(depth int, m *models.UserTeam, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveUserTeamAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveUserTeamIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveUserTeamAttributesFlags(depth int, m *models.UserTeam, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes UserTeamAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.UserTeamAttributes{}
	}

	err, attributesAdded := retrieveModelUserTeamAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveUserTeamIDFlags(depth int, m *models.UserTeam, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for UserTeamAttributes

// register flags to command
func registerModelUserTeamAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserTeamAttributesAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesBilling(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesCurrency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesOwner(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserTeamAttributesAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := ``

	var addressFlagName string
	if cmdPrefix == "" {
		addressFlagName = "address"
	} else {
		addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
	}

	var addressFlagDefault string

	_ = cmd.PersistentFlags().String(addressFlagName, addressFlagDefault, addressDescription)

	return nil
}

func registerUserTeamAttributesBilling(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var billingFlagName string
	if cmdPrefix == "" {
		billingFlagName = "billing"
	} else {
		billingFlagName = fmt.Sprintf("%v.billing", cmdPrefix)
	}

	if err := registerModelUserTeamAttributesBillingFlags(depth+1, billingFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserTeamAttributesCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdAtDescription := ``

	var createdAtFlagName string
	if cmdPrefix == "" {
		createdAtFlagName = "created_at"
	} else {
		createdAtFlagName = fmt.Sprintf("%v.created_at", cmdPrefix)
	}

	var createdAtFlagDefault string

	_ = cmd.PersistentFlags().String(createdAtFlagName, createdAtFlagDefault, createdAtDescription)

	return nil
}

func registerUserTeamAttributesCurrency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	currencyDescription := ``

	var currencyFlagName string
	if cmdPrefix == "" {
		currencyFlagName = "currency"
	} else {
		currencyFlagName = fmt.Sprintf("%v.currency", cmdPrefix)
	}

	var currencyFlagDefault string

	_ = cmd.PersistentFlags().String(currencyFlagName, currencyFlagDefault, currencyDescription)

	return nil
}

func registerUserTeamAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerUserTeamAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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

func registerUserTeamAttributesOwner(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ownerFlagName string
	if cmdPrefix == "" {
		ownerFlagName = "owner"
	} else {
		ownerFlagName = fmt.Sprintf("%v.owner", cmdPrefix)
	}

	if err := registerModelUserIncludeFlags(depth+1, ownerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserTeamAttributesSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	slugDescription := ``

	var slugFlagName string
	if cmdPrefix == "" {
		slugFlagName = "slug"
	} else {
		slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
	}

	var slugFlagDefault string

	_ = cmd.PersistentFlags().String(slugFlagName, slugFlagDefault, slugDescription)

	return nil
}

func registerUserTeamAttributesUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updatedAtDescription := ``

	var updatedAtFlagName string
	if cmdPrefix == "" {
		updatedAtFlagName = "updated_at"
	} else {
		updatedAtFlagName = fmt.Sprintf("%v.updated_at", cmdPrefix)
	}

	var updatedAtFlagDefault string

	_ = cmd.PersistentFlags().String(updatedAtFlagName, updatedAtFlagDefault, updatedAtDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserTeamAttributesFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveUserTeamAttributesAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, billingAdded := retrieveUserTeamAttributesBillingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || billingAdded

	err, createdAtAdded := retrieveUserTeamAttributesCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, currencyAdded := retrieveUserTeamAttributesCurrencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currencyAdded

	err, descriptionAdded := retrieveUserTeamAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrieveUserTeamAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, ownerAdded := retrieveUserTeamAttributesOwnerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ownerAdded

	err, slugAdded := retrieveUserTeamAttributesSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slugAdded

	err, updatedAtAdded := retrieveUserTeamAttributesUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	return nil, retAdded
}

func retrieveUserTeamAttributesAddressFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressFlagName := fmt.Sprintf("%v.address", cmdPrefix)
	if cmd.Flags().Changed(addressFlagName) {

		var addressFlagName string
		if cmdPrefix == "" {
			addressFlagName = "address"
		} else {
			addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
		}

		addressFlagValue, err := cmd.Flags().GetString(addressFlagName)
		if err != nil {
			return err, false
		}
		m.Address = &addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesBillingFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	billingFlagName := fmt.Sprintf("%v.billing", cmdPrefix)
	if cmd.Flags().Changed(billingFlagName) {
		// info: complex object billing UserTeamAttributesBilling is retrieved outside this Changed() block
	}
	billingFlagValue := m.Billing
	if swag.IsZero(billingFlagValue) {
		billingFlagValue = &models.UserTeamAttributesBilling{}
	}

	err, billingAdded := retrieveModelUserTeamAttributesBillingFlags(depth+1, billingFlagValue, billingFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || billingAdded
	if billingAdded {
		m.Billing = billingFlagValue
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesCreatedAtFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		createdAtFlagValue, err := cmd.Flags().GetString(createdAtFlagName)
		if err != nil {
			return err, false
		}
		m.CreatedAt = createdAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesCurrencyFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	currencyFlagName := fmt.Sprintf("%v.currency", cmdPrefix)
	if cmd.Flags().Changed(currencyFlagName) {

		var currencyFlagName string
		if cmdPrefix == "" {
			currencyFlagName = "currency"
		} else {
			currencyFlagName = fmt.Sprintf("%v.currency", cmdPrefix)
		}

		currencyFlagValue, err := cmd.Flags().GetString(currencyFlagName)
		if err != nil {
			return err, false
		}
		m.Currency = currencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesDescriptionFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = &descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesNameFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveUserTeamAttributesOwnerFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ownerFlagName := fmt.Sprintf("%v.owner", cmdPrefix)
	if cmd.Flags().Changed(ownerFlagName) {
		// info: complex object owner UserInclude is retrieved outside this Changed() block
	}
	ownerFlagValue := m.Owner
	if swag.IsZero(ownerFlagValue) {
		ownerFlagValue = &models.UserInclude{}
	}

	err, ownerAdded := retrieveModelUserIncludeFlags(depth+1, ownerFlagValue, ownerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ownerAdded
	if ownerAdded {
		m.Owner = ownerFlagValue
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesSlugFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	slugFlagName := fmt.Sprintf("%v.slug", cmdPrefix)
	if cmd.Flags().Changed(slugFlagName) {

		var slugFlagName string
		if cmdPrefix == "" {
			slugFlagName = "slug"
		} else {
			slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
		}

		slugFlagValue, err := cmd.Flags().GetString(slugFlagName)
		if err != nil {
			return err, false
		}
		m.Slug = slugFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesUpdatedAtFlags(depth int, m *models.UserTeamAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		updatedAtFlagValue, err := cmd.Flags().GetString(updatedAtFlagName)
		if err != nil {
			return err, false
		}
		m.UpdatedAt = updatedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for UserTeamAttributesBilling

// register flags to command
func registerModelUserTeamAttributesBillingFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserTeamAttributesBillingCustomerBillingID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamAttributesBillingID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserTeamAttributesBillingCustomerBillingID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	customerBillingIdDescription := ``

	var customerBillingIdFlagName string
	if cmdPrefix == "" {
		customerBillingIdFlagName = "customer_billing_id"
	} else {
		customerBillingIdFlagName = fmt.Sprintf("%v.customer_billing_id", cmdPrefix)
	}

	var customerBillingIdFlagDefault string

	_ = cmd.PersistentFlags().String(customerBillingIdFlagName, customerBillingIdFlagDefault, customerBillingIdDescription)

	return nil
}

func registerUserTeamAttributesBillingID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserTeamAttributesBillingFlags(depth int, m *models.UserTeamAttributesBilling, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, customerBillingIdAdded := retrieveUserTeamAttributesBillingCustomerBillingIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || customerBillingIdAdded

	err, idAdded := retrieveUserTeamAttributesBillingIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveUserTeamAttributesBillingCustomerBillingIDFlags(depth int, m *models.UserTeamAttributesBilling, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	customerBillingIdFlagName := fmt.Sprintf("%v.customer_billing_id", cmdPrefix)
	if cmd.Flags().Changed(customerBillingIdFlagName) {

		var customerBillingIdFlagName string
		if cmdPrefix == "" {
			customerBillingIdFlagName = "customer_billing_id"
		} else {
			customerBillingIdFlagName = fmt.Sprintf("%v.customer_billing_id", cmdPrefix)
		}

		customerBillingIdFlagValue, err := cmd.Flags().GetString(customerBillingIdFlagName)
		if err != nil {
			return err, false
		}
		m.CustomerBillingID = customerBillingIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserTeamAttributesBillingIDFlags(depth int, m *models.UserTeamAttributesBilling, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}
