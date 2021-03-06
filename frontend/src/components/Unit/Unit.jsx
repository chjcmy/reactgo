import React, { useCallback, useEffect, useRef, useState } from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import { Card, Image, Grid, Button, Label } from 'semantic-ui-react';
import { instance } from '../../axios';
import Server from '../../img/server.gif';
import Golang from '../../img/golang.gif';
import Js from '../../img/js.gif';
import Db from '../../img/db.gif';
import Life from '../../img/life.gif';
import Profile from '../../img/profile.jpg';

import './Unit.css';

const FamilyFont = styled.div`
  font-family: Neodgm, serif;
`;

const ImgCk = (unit) => {
  switch (unit) {
    case 1:
      return Golang;

    case 2:
      return Js;

    case 3:
      return Db;

    case 4:
      return Server;

    case 5:
      return Life;

    default:
      return null;
  }
};

// eslint-disable-next-line react/prop-types
const Unit = ({ match }) => {
  const isMounted = useRef(true);
  const [isSending, setIsSending] = useState(false);
  const [books, setBooks] = useState([]);

  const findAllBook = async () => {
    const res = await instance.get('/bookshow/0');
    setBooks(res.data);
  };
  // eslint-disable-next-line react-hooks/exhaustive-deps
  const findUnitBook = useCallback(async () => {
    const res = await instance.get(`/pickunitbooks/${match.params.unit}/0`);
    setBooks(res.data);
  }, [match.params.unit]);

  const deleteBook = useCallback(
    async (id) => {
      if (isSending) return;
      setIsSending(true);
      await instance.delete(`/bookdelete/${id}`);
      if (!match.params.unit) {
        findAllBook().then();
      } else {
        findUnitBook().then();
      }
      setIsSending(false);
    },
    [findUnitBook, isSending, match.params.unit]
  );

  useEffect(() => {
    isMounted.current = false;

    if (!match.params.unit) {
      findAllBook().then();
    } else {
      findUnitBook().then();
    }
    // eslint-disable-next-line react/prop-types
  }, [findUnitBook, match.params.unit]);

  useEffect(() => {}, []);

  return (
    <div>
      <FamilyFont>
        <Grid celled>
          {books.map((book) => (
            <Grid.Row key={book.id}>
              <Grid.Column width={5}>
                <Image
                  floated="left"
                  size="medium"
                  src={ImgCk(book.edges.unitid.id)}
                  style={{ position: 'inherit' }}
                />
              </Grid.Column>
              <Grid.Column width={11}>
                <Label image>
                  <img src={Profile} alt="profile" />
                  {book.edges.userid.name}
                </Label>
                <Card.Meta style={{ fontSize: 'large' }}>
                  ?????? ??????:
                  {book.create_at}
                  <Card.Meta style={{ fontSize: 'large' }}>
                    ??????????????? ??????:
                    {book.updated_at}
                  </Card.Meta>
                  <Card.Description style={{ fontSize: 'xxx-large' }}>
                    {book.title}
                  </Card.Description>
                  <Link to={`/book/${book.id}`}>
                    <Button inverted color="olive" size="big" floated="left">
                      ??????
                    </Button>
                  </Link>
                  {sessionStorage.getItem('id') ? (
                    <>
                      <Link to={`/bookupdate/${book.id}`}>
                        <Button
                          inverted
                          color="yellow"
                          size="big"
                          floated="left"
                        >
                          ????????????
                        </Button>
                      </Link>
                      <Button
                        inverted
                        color="red"
                        size="big"
                        floated="left"
                        onClick={() => deleteBook(book.id)}
                      >
                        ??????
                      </Button>
                    </>
                  ) : null}
                </Card.Meta>
              </Grid.Column>
            </Grid.Row>
          ))}
        </Grid>
      </FamilyFont>
    </div>
  );
};

export default Unit;
