import { useEffect } from 'react';
import { CharacterSelectView } from './CharacterSelect.view';

type CharacterSelectControllerProps = {
  selectedCharacter: string;
  setSelectedCharacter: (character: string) => void;
}

export const CharacterSelectController = ({ selectedCharacter, setSelectedCharacter }: CharacterSelectControllerProps) => {
  // @TODO: Replace this with a call to the server to get the list of characters
  const characters = ['Maunders', 'Someone Else']

  useEffect(() => {
    console.log('Setting selected character to first character in list')
    setSelectedCharacter(characters[0]);
  }, []);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSelectedCharacter(event.target.value);
  };

  return (
    <CharacterSelectView
      characters={characters}
      selectedCharacter={selectedCharacter}
      handleChange={handleChange}
    />
  );
};
