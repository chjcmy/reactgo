import React, {useEffect, useRef, useState} from 'react';
import {instance} from "../../axios";
import styled from "styled-components";
import {makeStyles} from '@material-ui/core/styles';
import {Button, Card, Image} from 'semantic-ui-react'

import Grid from '@material-ui/core/Grid';

const FamilyFont = styled.div`
  font-family: Neodgm, serif;
`;

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        padding: theme.spacing(1),
        textAlign: 'center',
        color: theme.palette.text.secondary,
    },
}));

const ImgCk = (unitid) => {
    switch (unitid) {
        case 1 :
            return "https://media1.giphy.com/media/MCRLiCTdFtKYsF1xOb/giphy.gif?cid=ecf05e47m17bngd5zk6tk6hxcv6zublos400em0wnujy2va1&rid=giphy.gif&ct=g"

        case 2 :
            return "https://media0.giphy.com/media/fuJPZBIIqzbt1kAYVc/giphy.gif?cid=ecf05e47a6aka78ap4znwlsxkhsssw3n08ad1413soxbsc5z&rid=giphy.gif&ct=g"

        case 3 :
            return "https://media3.giphy.com/media/XsHkc4MCBXDn0yNybG/giphy.gif?cid=ecf05e47t53tpe7oz4fdbwe2f9wlp8nndofn0hds0adjlus0&rid=giphy.gif&ct=g"

        case 4 :
            return "https://media0.giphy.com/media/GbH8vRmrNHdVZhouBt/giphy.gif?cid=790b761125661d20b636c1f65db78b00891724e42a74254a&rid=giphy.gif&ct=g"

        case 5 :
            return "https://media1.giphy.com/media/17XAjPucc8Qda/giphy.gif?cid=ecf05e47acbjnzaliu1tjp1prqkmbwddng7zwd1aeyr1xjrm&rid=giphy.gif&ct=g"
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
                                    <div className="nes-container with-title" style={{position: "inherit", margin: "5%"}}>
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
                                            <div className='ui two buttons'>
                                                <Button basic color='green'>
                                                    Approve
                                                </Button>
                                                <Button basic color='red'>
                                                    Decline
                                                </Button>
                                            </div>
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
