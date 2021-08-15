import React, {useEffect, useState} from 'react';
import "nes.css/css/nes.min.css";
import {instance} from "../../../axios";
import 'suneditor/dist/css/suneditor.min.css';
import SunEditor from "suneditor-react";
import styled from "styled-components";

const BookDiv= styled.div`
  margin: 2% 2% 2% 2%;
  width: 100%;
  display: block;
`

const BookText = styled.div`
  margin: auto;
  text-align: center;
  align-items: center;
  justify-content: flex-start;
  display: flex;
  height: 50%;
`

const Date = styled.div`
  font-size: large;
`

const Book = ({match}: { match: any }) => {

    console.log(match.params.id)

    const [rbook, setRbook] = useState<any>([]);

    const findBook = async () => {
        const res = await instance.get(`/bookread/${match.params.id}`)
        setRbook(res.data)
        console.log(rbook)
    };

    useEffect(() => {
        findBook().then(() => console.log(rbook));
    }, [match.params.unit]);

    return (
        <>
            <BookDiv>
            <div style={{fontFamily:"Neodgm", fontSize:"xx-large"}}>주제: {rbook.title}</div>
            <div style={{fontFamily:"Neodgm", }}>
                <div style={{fontSize:"x-large"}}>글쓴이: {rbook.edges?.userid.name}</div>
                <div style={{fontSize:"x-large"}}>컨텐츠: {rbook.edges?.unitid.content_name}</div>
                <Date>만든 날짜: {rbook.create_at}</Date>
                <Date>업데이트한 날짜: {rbook.updated_at}</Date>
                <SunEditor hideToolbar={true} setContents={rbook.subject} height="25rem"/>
            </div>
            </BookDiv>
        </>
    );
};

export default Book;
