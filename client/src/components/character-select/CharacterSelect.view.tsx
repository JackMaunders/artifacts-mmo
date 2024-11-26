import {
  Box,
  FormControl,
  FormControlLabel,
  FormLabel,
  Radio,
  RadioGroup,
  useTheme
} from "@mui/material";

type CharacterSelectViewProps = {
  characters: string[];
  selectedCharacter: string;
  handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

export const CharacterSelectView = ({ characters, selectedCharacter, handleChange }: CharacterSelectViewProps) => {
  const theme = useTheme();

  return (
    <Box sx={{ marginTop: theme.spacing(2) }}>
      <FormControl>
        <FormLabel id="character-select-radio-buttons-group">Character</FormLabel>
        <RadioGroup
          aria-labelledby="character-select-controlled-radio-buttons-group"
          name="controlled-radio-buttons-group"
          value={selectedCharacter}
          onChange={handleChange}
        >
          {characters.map((character) => (
            <FormControlLabel
              key={character}
              value={character}
              control={<Radio />}
              label={character}
            />
          ))}
        </RadioGroup>
      </FormControl>
    </Box>
  );
};
