import React, {useEffect, useRef, useState} from 'react';
import {instance} from "../../axios";
import styled from "styled-components";
import {makeStyles} from '@material-ui/core/styles';
import {Button, Card, Image} from 'semantic-ui-react'
import Server from '../../img/server.gif'
import Golang from '../../img/golang.gif'
import Js from '../../img/js.gif'
import Db from '../../img/db.gif'
import Life from '../../img/life.gif'


import Grid from '@material-ui/core/Grid';

const FamilyFont = styled.div`
  font-family: Neodgm, serif;
`;

const ImgCk = (unitid) => {
    switch (unitid) {
        case 1 :
            return Golang

        case 2 :
            return Js

        case 3 :
            return Db

        case 4 :
            return Server

        case 5 :
            return Life
    }
}


const Unit = ({match}) => {

    useRef()

    const [login, setLogin] = useState(!!localStorage.getItem('id'));
    const [books, setBooks] = useState([]);

    const findAllBook = async () => {
        const res = await instance.get('/bookshow/0');
        setBooks(res.data);
        console.log(res.data);
    };
    const findUnitBook = async () => {
        const res = await instance.get(`/pickunitbooks/${match.params.unit}/0`);
        setBooks(res.data);
        console.log(books);
    };
    useEffect(() => {
        if (!match.params.unit) {
            findAllBook().then();
        } else {
            findUnitBook().then();
        }
    }, [match.params.unit]);

    console.log(books);
    return (
        <div>
            <FamilyFont>
                <Grid container spacing={1}>
                    <Grid container direction="row"
                          justifyContent="center"
                          alignItems="center">
                        {books.map((book, idx) => (
                                <Grid item xs={6}>
                                    <div id={idx} className="nes-container with-title" style={{position: "inherit", margin: "5%"}}>
                                        <p className="title"
                                           style={{fontSize: "xx-large"}}>{book.edges.unitid.content_name}</p>
                                        <Card.Content>
                                            <Image
                                                floated='right'
                                                size='small'
                                                src={ImgCk(book.edges.unitid.id)}
                                                style={{position: "inherit"}}
                                            />
                                            <Card.Meta style={{fontSize: "large"}}>만든 날짜</Card.Meta>
                                            <Card.Meta style={{fontSize: "large"}}>{book.create_at}</Card.Meta>
                                            <Card.Meta style={{fontSize: "large"}}>업데이트된 날짜</Card.Meta>
                                            <Card.Meta style={{fontSize: "large"}}>{book.updated_at}</Card.Meta>
                                            <Card.Meta
                                                style={{fontSize: "x-large"}}>글쓴이: {book.edges.userid.name}</Card.Meta>
                                            <Card.Description style={{fontSize: "xx-large"}}>
                                                {book.title}
                                            </Card.Description>
                                        </Card.Content>
                                        <Card.Content extra>
                                        </Card.Content>
                                    </div>
                                </Grid>
                            )
                        )}
                    </Grid>
                </Grid>
            </FamilyFont>
        </div>
    );
};

export default Unit;
