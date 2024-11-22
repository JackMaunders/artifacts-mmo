import './Gamepad.css'

export const GamepadView = () => {
  return (
    <div className='controllerContainer'>
      <div className="cable" />
      <div className="controller">
        <div className="centerBlue">
          <div className="centerLeft" />
          <div className="centerRight" />
        </div>
        <div className="centerStart">
          <div className="SLeft" />
          <div className="SRight" />
        </div>
        <div className="centerSelect">
          <div className="SLeft" />
          <div className="SRight" />
        </div>
        <div className="controllerLeft">
          <div className="circle" />
          <div className="crossCenter">
            <div className="crossTop" />
            <div className="crossBottom" />
            <div className="crossLeft" />
            <div className="crossRight" />
            <div className="crossCircle" />
          </div>
        </div>
        <div className="controllerRight">
          <div className="backButton1Center">
            <div className="cornerLeft1" />
            <div className="cornerRight1" />
          </div>
          <div className="backButton2Center">
            <div className="cornerLeft2" />
            <div className="cornerRight2" />
          </div>
        </div>
      </div>
    </div>
  );
};
