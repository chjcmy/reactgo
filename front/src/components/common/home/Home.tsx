import React from 'react';
import styled from "styled-components";

const NewBooks = styled.div`
  text-align: center;
  margin-left: 45%;
  margin-top: 2%;
  font-family: Neodgm, serif;
`;

const Home = () => {
    return (
            <>
                <NewBooks>
                    <h1>최신글</h1>
                </NewBooks>

            </>
    );
};

export default Home;
