import React, {useEffect, useState} from 'react';
import "nes.css/css/nes.min.css";
import {instance} from "../../../axios";



const Book = ({match} : {match: any}) => {

    console.log(match.params.id)

    const [book, setBook] = useState<any[]>([]);

    const findBook = async () => {
        const res = await instance.get(`/bookread/${match.params.id}`)
        setBook(res.data)
        console.log(book)
    };

    useEffect(() => {
            findBook()
        console.log(book)
    }, [match.params.unit]);

    useEffect(() => {
    }, [match.params.unit]);

    return (
        <>
                <h1>Book</h1>
        </>
    );
};

export default Book;
