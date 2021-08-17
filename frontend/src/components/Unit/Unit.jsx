import React, { useEffect, useState } from 'react';
import { instance } from "../../axios";

const Unit = ({ match }) => {
    const [books, setBooks] = useState([]);
    const findAllBook = async () => {
        const res = await instance.get('/bookshow/0');
        setBooks(res.data);
        console.log(res.data);
    };
    const findUnitBook = async () => {
        const res = await instance.get(`/pickunitbooks/${match.params.unit}/1`);
        setBooks(res.data);
        console.log(books);
    };
    useEffect(() => {
        if (!match.params.unit) {
            findAllBook().then();
        }
        else {
            findUnitBook().then();
        }
    }, [match.params.unit]);
    return (
        <div>
            {match.params.unit}
            {books.map(books => (
                    <div key={books.id}>
                        {books.subject}
                    </div>
                )
            )}
        </div>
    );
};

export default Unit;