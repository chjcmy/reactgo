import React, { useEffect, useState } from 'react';
import styled from "styled-components";
import "nes.css/css/nes.min.css";
import {Link} from 'react-router-dom'
import {instance} from "../../axios";
import GoogleAdsense from "react-adsense-google";
const NewBooks = styled.div `
  margin-top: 5%;
  text-align: center;
  justify-content: flex-start, !important;
  align-items: center;
  font-size: 80%;
  font-family: Neodgm, serif;
`;
const Home = () => {
    const [newBooks, setNewBooks] = useState([]);
    const findNewBooks = async () => {
        const res = await instance.get('/newbooks');
        setNewBooks(res.data);
    };

    useEffect(() => {
        findNewBooks().then();
    }, []);
    return (
        <>
            <NewBooks>
                <h1>최신글</h1>
                {newBooks.map(newbook => (
                    <div className="nes-container with-title is-centered"
                         style={{position: "inherit", marginBlock: "5%"}} key={newbook.id}>
                        <p className="title title-font" style={{fontSize: "x-large"}}>
                            {newbook.edges.unitid.content_name}
                        </p>
                        <Link to={`/book/${newbook.id}`}>{newbook.title}</Link>
                    </div>
                ))}
            </NewBooks>
        </>
    );
};
export default Home;
