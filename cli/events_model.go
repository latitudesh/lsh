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

// Schema cli for Events

// register flags to command
func registerModelEventsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEventsAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelEventsAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelEventsFlags(depth int, m *models.Events, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveEventsAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveEventsIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveEventsAttributesFlags(depth int, m *models.Events, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes EventsAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.EventsAttributes{}
	}

	err, attributesAdded := retrieveModelEventsAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveEventsIDFlags(depth int, m *models.Events, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for EventsAttributes

// register flags to command
func registerModelEventsAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEventsAttributesAction(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesAuthor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesProject(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesTarget(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesTeam(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesAction(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	actionDescription := ``

	var actionFlagName string
	if cmdPrefix == "" {
		actionFlagName = "action"
	} else {
		actionFlagName = fmt.Sprintf("%v.action", cmdPrefix)
	}

	var actionFlagDefault string

	_ = cmd.PersistentFlags().String(actionFlagName, actionFlagDefault, actionDescription)

	return nil
}

func registerEventsAttributesAuthor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var authorFlagName string
	if cmdPrefix == "" {
		authorFlagName = "author"
	} else {
		authorFlagName = fmt.Sprintf("%v.author", cmdPrefix)
	}

	if err := registerModelEventsAttributesAuthorFlags(depth+1, authorFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEventsAttributesProject(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var projectFlagName string
	if cmdPrefix == "" {
		projectFlagName = "project"
	} else {
		projectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
	}

	if err := registerModelEventsAttributesProjectFlags(depth+1, projectFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesTarget(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var targetFlagName string
	if cmdPrefix == "" {
		targetFlagName = "target"
	} else {
		targetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
	}

	if err := registerModelEventsAttributesTargetFlags(depth+1, targetFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesTeam(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var teamFlagName string
	if cmdPrefix == "" {
		teamFlagName = "team"
	} else {
		teamFlagName = fmt.Sprintf("%v.team", cmdPrefix)
	}

	if err := registerModelEventsAttributesTeamFlags(depth+1, teamFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEventsAttributesFlags(depth int, m *models.EventsAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, actionAdded := retrieveEventsAttributesActionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || actionAdded

	err, authorAdded := retrieveEventsAttributesAuthorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorAdded

	err, createdAtAdded := retrieveEventsAttributesCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, projectAdded := retrieveEventsAttributesProjectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectAdded

	err, targetAdded := retrieveEventsAttributesTargetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded

	err, teamAdded := retrieveEventsAttributesTeamFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || teamAdded

	return nil, retAdded
}

func retrieveEventsAttributesActionFlags(depth int, m *models.EventsAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	actionFlagName := fmt.Sprintf("%v.action", cmdPrefix)
	if cmd.Flags().Changed(actionFlagName) {

		var actionFlagName string
		if cmdPrefix == "" {
			actionFlagName = "action"
		} else {
			actionFlagName = fmt.Sprintf("%v.action", cmdPrefix)
		}

		actionFlagValue, err := cmd.Flags().GetString(actionFlagName)
		if err != nil {
			return err, false
		}
		m.Action = actionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEventsAttributesAuthorFlags(depth int, m *models.EventsAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authorFlagName := fmt.Sprintf("%v.author", cmdPrefix)
	if cmd.Flags().Changed(authorFlagName) {
		// info: complex object author EventsAttributesAuthor is retrieved outside this Changed() block
	}
	authorFlagValue := m.Author
	if swag.IsZero(authorFlagValue) {
		authorFlagValue = &models.EventsAttributesAuthor{}
	}

	err, authorAdded := retrieveModelEventsAttributesAuthorFlags(depth+1, authorFlagValue, authorFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorAdded
	if authorAdded {
		m.Author = authorFlagValue
	}

	return nil, retAdded
}

func retrieveEventsAttributesCreatedAtFlags(depth int, m *models.EventsAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEventsAttributesProjectFlags(depth int, m *models.EventsAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectFlagName := fmt.Sprintf("%v.project", cmdPrefix)
	if cmd.Flags().Changed(projectFlagName) {
		// info: complex object project EventsAttributesProject is retrieved outside this Changed() block
	}
	projectFlagValue := m.Project
	if swag.IsZero(projectFlagValue) {
		projectFlagValue = &models.EventsAttributesProject{}
	}

	err, projectAdded := retrieveModelEventsAttributesProjectFlags(depth+1, projectFlagValue, projectFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectAdded
	if projectAdded {
		m.Project = projectFlagValue
	}

	return nil, retAdded
}

func retrieveEventsAttributesTargetFlags(depth int, m *models.EventsAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	targetFlagName := fmt.Sprintf("%v.target", cmdPrefix)
	if cmd.Flags().Changed(targetFlagName) {
		// info: complex object target EventsAttributesTarget is retrieved outside this Changed() block
	}
	targetFlagValue := m.Target
	if swag.IsZero(targetFlagValue) {
		targetFlagValue = &models.EventsAttributesTarget{}
	}

	err, targetAdded := retrieveModelEventsAttributesTargetFlags(depth+1, targetFlagValue, targetFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded
	if targetAdded {
		m.Target = targetFlagValue
	}

	return nil, retAdded
}

func retrieveEventsAttributesTeamFlags(depth int, m *models.EventsAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	teamFlagName := fmt.Sprintf("%v.team", cmdPrefix)
	if cmd.Flags().Changed(teamFlagName) {
		// info: complex object team EventsAttributesTeam is retrieved outside this Changed() block
	}
	teamFlagValue := m.Team
	if swag.IsZero(teamFlagValue) {
		teamFlagValue = &models.EventsAttributesTeam{}
	}

	err, teamAdded := retrieveModelEventsAttributesTeamFlags(depth+1, teamFlagValue, teamFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || teamAdded
	if teamAdded {
		m.Team = teamFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for EventsAttributesAuthor

// register flags to command
func registerModelEventsAttributesAuthorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEventsAttributesAuthorEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesAuthorID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesAuthorName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesAuthorEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEventsAttributesAuthorID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEventsAttributesAuthorName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEventsAttributesAuthorFlags(depth int, m *models.EventsAttributesAuthor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, emailAdded := retrieveEventsAttributesAuthorEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, idAdded := retrieveEventsAttributesAuthorIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveEventsAttributesAuthorNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveEventsAttributesAuthorEmailFlags(depth int, m *models.EventsAttributesAuthor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEventsAttributesAuthorIDFlags(depth int, m *models.EventsAttributesAuthor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEventsAttributesAuthorNameFlags(depth int, m *models.EventsAttributesAuthor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for EventsAttributesProject

// register flags to command
func registerModelEventsAttributesProjectFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEventsAttributesProjectID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesProjectName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesProjectSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesProjectID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEventsAttributesProjectName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEventsAttributesProjectSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEventsAttributesProjectFlags(depth int, m *models.EventsAttributesProject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveEventsAttributesProjectIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveEventsAttributesProjectNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, slugAdded := retrieveEventsAttributesProjectSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slugAdded

	return nil, retAdded
}

func retrieveEventsAttributesProjectIDFlags(depth int, m *models.EventsAttributesProject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEventsAttributesProjectNameFlags(depth int, m *models.EventsAttributesProject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEventsAttributesProjectSlugFlags(depth int, m *models.EventsAttributesProject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for EventsAttributesTarget

// register flags to command
func registerModelEventsAttributesTargetFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEventsAttributesTargetID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesTargetName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesTargetID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEventsAttributesTargetName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEventsAttributesTargetFlags(depth int, m *models.EventsAttributesTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveEventsAttributesTargetIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveEventsAttributesTargetNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveEventsAttributesTargetIDFlags(depth int, m *models.EventsAttributesTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEventsAttributesTargetNameFlags(depth int, m *models.EventsAttributesTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for EventsAttributesTeam

// register flags to command
func registerModelEventsAttributesTeamFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEventsAttributesTeamID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEventsAttributesTeamName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEventsAttributesTeamID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEventsAttributesTeamName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEventsAttributesTeamFlags(depth int, m *models.EventsAttributesTeam, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveEventsAttributesTeamIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveEventsAttributesTeamNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveEventsAttributesTeamIDFlags(depth int, m *models.EventsAttributesTeam, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEventsAttributesTeamNameFlags(depth int, m *models.EventsAttributesTeam, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
