import { Container, createTheme, CssBaseline, Paper, ThemeProvider, Typography } from '@mui/material'
import { CharacterSelect } from './components/character-select'
import { Gamepad } from './components/game-pad'
import { useState } from 'react';

const theme = createTheme({
  typography: {
    fontSize: 14,
  },
});

function App() {
  const [selectedCharacter, setSelectedCharacter] = useState('')

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Container>
          <Gamepad selectedCharacter={selectedCharacter} />
          <Paper sx={{ padding: theme.spacing(4) }}>
            <Typography variant="h3" component="h1">Artifacts MMO Controller</Typography>
            <Typography variant="body1" component="p">1. Select a character to control</Typography>
            <Typography variant="body1" component="p">2. Use the gamepad to move the character</Typography>
            <CharacterSelect selectedCharacter={selectedCharacter} setSelectedCharacter={setSelectedCharacter} />
          </Paper>
      </Container>
    </ThemeProvider>
  )
}

export default App
