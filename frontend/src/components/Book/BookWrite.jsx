import React, {useEffect, useState} from 'react';
import {instance} from '../../axios';
import styled from 'styled-components';
import {Link} from 'react-router-dom';
import {FormControl, InputLabel, makeStyles, MenuItem, Select} from "@material-ui/core";
import Editor from 'ckeditor5-custom-build/build/ckeditor';
import { CKEditor } from '@ckeditor/ckeditor5-react'
import ClassicEditor from "@ckeditor/ckeditor5-editor-classic/src/classiceditor";
import Bold from "@ckeditor/ckeditor5-basic-styles/src/bold";
import Italic from "@ckeditor/ckeditor5-basic-styles/src/italic";
import Essentials from "@ckeditor/ckeditor5-essentials/src/essentials";
import Paragraph from "@ckeditor/ckeditor5-paragraph/src/paragraph";

const editorConfiguration = {
    plugins: [Paragraph, Bold, Italic, Essentials],
    toolbar: ["bold", "italic"],
};

const Outer = styled.div`
  margin: 2% 1% 0 2%;
  width: 80%;
`;

const FamilyFont = styled.div`
  font-family: Neodgm, serif;
`;

const Creator = styled.div`
  display: flex;
  height: 50%;
`;
const Save = styled(Link)`
  display: flex;
  height: 50%;
`;


const
useStyles = makeStyles((theme) => ({
    button: {
        display: 'block',
        marginTop: theme.spacing(2),
    },
    formControl: {
        margin: theme.spacing(1),
        minWidth: 120,
    },
}));
const BookWrite = () => {

    const classes = useStyles();
    const [units, setUnits] = useState([]);
    const [read, setRead] = useState(false);
    const [chapter, setChapter] = useState('');
    const [title, setTitle] = useState('');
    const [subject, setSubject] = useState('');
    const [open, setOpen] = React.useState(false);
    const findUnits = async () => {
        await instance.get('/unitshosting').then(function (res) {
            setUnits(res.data);
            console.log(units);
        })
            .catch(function (error) {
                console.log(error);
            });
    };
    const createBook = async () => {
        console.log(chapter, title, subject);
        await instance.post('/bookcreate', {
            unit: chapter,
            title: title,
            subject: subject,
        }, {
            headers: {
                'Content-Type': 'application/json',
                'id': localStorage.getItem("id")
            }
        });
    };
    if (read) {
        setTimeout(() => setRead(false), 500);
    }
    useEffect(() => {
        findUnits().then();
    }, []);
    const click = () => {
        console.log(chapter, title, subject);
        createBook().then();
    };
    const handleClose = () => {
        setOpen(false);
    };
    const handleOpen = () => {
        setOpen(true);
    };
    const handleChange = (event) => {
        setChapter(event.target.value);
    };

    return <>
        <Outer>
            <FamilyFont>
                <Creator>
                    <div>
                        <FormControl style={{width: "100%"}} className={classes.formControl}>
                            <InputLabel id="demo-controlled-open-select-label">unit</InputLabel>
                            <Select
                                labelId="demo-controlled-open-select-label"
                                id="demo-controlled-open-select"
                                open={open}
                                onClose={handleClose}
                                onOpen={handleOpen}
                                value={chapter}
                                onChange={handleChange}
                            >
                                <MenuItem value="">
                                    <em>None</em>
                                </MenuItem>
                                {units.map(unit => <MenuItem value={unit.id}>{unit.content_name}</MenuItem>
                                )}
                            </Select>
                        </FormControl>
                    </div>
                    {read ?
                        <span className="nes-text is-success" style={{
                            fontSize: "xxx-large",
                            marginRight: "2%",
                            marginTop: "2%",
                            marginLeft: "3%"
                        }}>읽혀짐</span>
                        :
                        null
                    }
                    <Save to="/" onClick={click} style={{marginLeft: "30%", width: "100%"}}>
                        <button type="button" className="nes-btn is-primary"
                                style={{marginLeft: "2%", fontSize: "xx-large"}}>저장
                        </button>
                    </Save>
                </Creator>
                <div className="nes-field" style={{marginBottom: "2%"}}>
                    <label htmlFor="name_field">주제</label>
                    <input type="text" id="name_field" className="nes-input" style={{width: "90%", height: "50%"}}
                           onChange={e => setTitle(e.target.value)}/>
                </div>
            </FamilyFont>
            <CKEditor
                editor={ClassicEditor}
                data='<p>Hello from CKEditor 5!</p>'
                config={editorConfiguration}
                onChange={(event, editor) => {
                    const data = editor.getData();
                    console.log(data);
                }}
            />
        </Outer>
    </>;
};

export default BookWrite;
