import React, {useEffect,  useState} from 'react';
import {Link} from 'react-router-dom'
import {instance} from "../../axios";
import styled from "styled-components";
import {Card, Image, Grid, Button, Label} from 'semantic-ui-react'
import Server from '../../img/server.gif'
import Golang from '../../img/golang.gif'
import Js from '../../img/js.gif'
import Db from '../../img/db.gif'
import Life from '../../img/life.gif'
import Profile from '../../img/profile.jpg'

import './Unit.css'

const FamilyFont = styled.div`
  font-family: Neodgm, serif;
`;

const ImgCk = (unit) => {
    switch (unit) {
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

        default :
            return null
    }
}

const Unit = ({match}) => {

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
                <Grid celled>
                        {books.map((book, idx) => (
                            <Grid.Row key={idx}>
                                <Grid.Column width={5}>
                                    <Image
                                        floated='left'
                                        size='medium'
                                        src={ImgCk(book.edges.unitid.id)}
                                        style={{position: "inherit"}}
                                    />
                                </Grid.Column>
                                <Grid.Column width={11}>
                                    <Label image>
                                        <img src={Profile} />
                                        {book.edges.userid.name}
                                    </Label>
                                            <Card.Meta style={{fontSize: "large"}}>만든 날짜: {book.create_at}</Card.Meta>
                                            <Card.Meta style={{fontSize: "large"}}>업데이트된 날짜: {book.updated_at}</Card.Meta>
                                            <Card.Description style={{fontSize: "xxx-large"}}>
                                                {book.title}
                                            </Card.Description>
                                    <Link to={`/book/${book.id}`}>
                                    <Button inverted color='olive' size='big' floated={"left"}>
                                        보기
                                    </Button>
                                    </Link>
                                    { sessionStorage.getItem('id') === '1' ?
                                        <>
                                        <Button inverted color='yellow' size='big' floated={"left"}>
                                            업데이트
                                        </Button>
                                        <Button inverted color='red' size='big' floated={"left"}>
                                        삭제
                                        </Button>
                                        </>
                                        :
                                        null
                                    }
                                </Grid.Column>
                            </Grid.Row>
                            )
                        )}
                </Grid>
            </FamilyFont>
        </div>
    );
};

export default Unit;
