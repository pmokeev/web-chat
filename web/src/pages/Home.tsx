import React from 'react';

const Home = (props: { isJWTCorrect: boolean }) => {
  return (
    <div>
      {props.isJWTCorrect ? 'Hi!' : 'You are not logged in'}
    </div>
  );
};

export default Home;