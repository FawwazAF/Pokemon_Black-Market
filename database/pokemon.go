package database

import (
	"project/pokemon/config"
	"project/pokemon/models"
)

func AddPokemon(pokemon models.Pokemon) (models.Pokemon, error) {
	if err := config.DB.Save(&pokemon).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func GetPokemonFromDatabase(pokemon_id int) (models.Pokemon, error) {
	var pokemon models.Pokemon
	if err := config.DB.Find(&pokemon, "id=?", pokemon_id).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func EditPokemon(pokemon models.Pokemon) (models.Pokemon, error) {
	if err := config.DB.Save(&pokemon).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func DeletePokemon(pokemon_id int) ([]models.Pokemon, error) {
	var pokemon []models.Pokemon
	if err := config.DB.Find(&pokemon, "id=?", pokemon_id).Error; err != nil {
		return pokemon, err
	}
	if err := config.DB.Delete(&pokemon, "id=?", pokemon_id).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func SearchPokemon(pokemon_name string) ([]models.Pokemon, error) {
	var pokemon []models.Pokemon
	search_key := ("%" + pokemon_name + "%")
	if err := config.DB.Find(&pokemon, "name LIKE ?", search_key).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}
