import { GamepadView } from './Gamepad.view';

export const GamepadController = ({ selectedCharacter }: { selectedCharacter: string }) => {
  const handleDirection = async (direction: string) => {
    // @TODO: Move the sending of requests to backend to a service & add error handling
    console.log(`Moving ${selectedCharacter} ${direction}`);

    try {
      await fetch(`http://localhost:8080/${selectedCharacter}/move/${direction}`, { method: 'POST' });
    } catch (error) {
      console.error('Error moving character', error);
    }
  };

  return <GamepadView handleDirection={handleDirection} />;
};
