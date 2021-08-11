import React, {FC, useEffect, useState} from 'react';
import styled from "styled-components";
import {instance} from "../../../axios";
import "nes.css/css/nes.min.css";
import "./Home.css"
import {Link} from "@material-ui/core";

const NewBooks = styled.div`
  margin: auto;
  text-align: center;
  align-items: center;
  font-size: 80%;
  font-family: Neodgm, serif;
`;

const Home: FC = () => {

    const [newBooks, setNewBooks] = useState<any[]>([]);

    const findNewBooks = async () => {
        const res = await instance.get('/newbooks')
            .catch(function (error: any) {
                    console.log(error)
                }
            );
        setNewBooks(res.data)
        console.log(newBooks)
    };

    useEffect(() => {
        findNewBooks().then(r => console.log(r));
    }, []);

    return (
        <>
            <NewBooks>
                <h1>최신글</h1>
                {newBooks.map(newbook => (
                    <div className="nes-container with-title is-centered"
                         style={{position: "inherit", marginBlock: "20%"}} key={newbook.id}>
                        <p className="title title-font" style={{fontSize: "x-large"}}>
                            {newbook.edges.unitid.content_name}
                        </p>
                        <Link href={`/book/${newbook.id}`}>{newbook.subject}</Link>
                    </div>
                ))}
            </NewBooks>
        </>
    );
};

export default Home;
