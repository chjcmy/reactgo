import React, { useEffect, useState } from 'react';
import "nes.css/css/nes.min.css";
import { instance } from "../../axios";
import styled from "styled-components";
import "./ckcontent.css"
import { CKEditor } from '@ckeditor/ckeditor5-react';

import ClassicEditor from '@ckeditor/ckeditor5-editor-classic/src/classiceditor';
import Bold from '@ckeditor/ckeditor5-basic-styles/src/bold';
import Italic from '@ckeditor/ckeditor5-basic-styles/src/italic';
import Essentials from '@ckeditor/ckeditor5-essentials/src/essentials';
import Paragraph from '@ckeditor/ckeditor5-paragraph/src/paragraph';
import Heading from "@ckeditor/ckeditor5-heading/src/heading";
import Underline from "@ckeditor/ckeditor5-basic-styles/src/underline";
import Strikethrough from "@ckeditor/ckeditor5-basic-styles/src/strikethrough";
import BlockQuote from "@ckeditor/ckeditor5-block-quote/src/blockquote";
import MediaEmbed from "@ckeditor/ckeditor5-media-embed/src/mediaembed";
import PasteFromOffice from "@ckeditor/ckeditor5-paste-from-office/src/pastefromoffice";
import Font from "@ckeditor/ckeditor5-font/src/font";
import Image from "@ckeditor/ckeditor5-image/src/image";
import ImageStyle from "@ckeditor/ckeditor5-image/src/imagestyle";
import ImageToolbar from "@ckeditor/ckeditor5-image/src/imagetoolbar";
import ImageUpload from "@ckeditor/ckeditor5-image/src/imageupload";
import ImageResize from "@ckeditor/ckeditor5-image/src/imageresize";
import List from "@ckeditor/ckeditor5-list/src/list";
import Alignment from "@ckeditor/ckeditor5-alignment/src/alignment";
import Table from "@ckeditor/ckeditor5-table/src/table";
import TableToolbar from "@ckeditor/ckeditor5-table/src/tabletoolbar";
import TextTransformation from "@ckeditor/ckeditor5-typing/src/texttransformation";
import Indent from "@ckeditor/ckeditor5-indent/src/indent";
import IndentBlock from "@ckeditor/ckeditor5-indent/src/indentblock";
import Base64UploadAdapter from "@ckeditor/ckeditor5-upload/src/adapters/base64uploadadapter";
import CodeBlock from "@ckeditor/ckeditor5-code-block/src/codeblock";

const editorConfiguration = {
    language: "ko",
    plugins: [
        Essentials,
        Paragraph,
        Bold,
        Italic,
        Heading,
        Underline,
        Strikethrough,
        BlockQuote,
        MediaEmbed,
        PasteFromOffice,
        Font,
        Image,
        ImageStyle,
        ImageToolbar,
        ImageUpload,
        ImageResize,
        List,
        Alignment,
        Table,
        TableToolbar,
        TextTransformation,
        Indent,
        IndentBlock,
        Base64UploadAdapter,
        CodeBlock,
    ],
    alignment: {
        options: ["justify", "left", "center", "right"],
    },
    table: {
        contentToolbar: ["tableColumn", "tableRow", "mergeTableCells"],
    },
    image: {
        resizeUnit: "px",
        toolbar: [
            "imageStyle:alignLeft",
            "imageStyle:full",
            "imageStyle:alignRight",
            "|",
            "imageTextAlternative",
        ],
        styles: ["full", "alignLeft", "alignRight"],
    },
    typing: {
        transformations: {
            remove: [
                "enDash",
                "emDash",
                "oneHalf",
                "oneThird",
                "twoThirds",
                "oneForth",
                "threeQuarters",
            ],
        },
    },
    codeBlock: {
        languages: [
            { language: 'plaintext', label: 'Plain text' },
            { language: 'c', label: 'C' },
            { language: 'cs', label: 'C#' },
            { language: 'cpp', label: 'C++' },
            { language: 'css', label: 'CSS' },
            { language: 'diff', label: 'Diff' },
            { language: 'html', label: 'HTML' },
            { language: 'java', label: 'Java' },
            { language: 'javascript', label: 'JavaScript' },
            { language: 'php', label: 'PHP' },
            { language: 'python', label: 'Python' },
            { language: 'ruby', label: 'Ruby' },
            { language: 'typescript', label: 'TypeScript' },
            { language: 'xml', label: 'XML' },
            { language: 'go', label: 'Go'}
        ],
    },
}

const BookDiv = styled.div `
  margin: 2% 2% 2% 2%;
  width: 100%;
  display: block;
`;
const BookText = styled.div `
  margin: auto;
  text-align: center;
  align-items: center;
  justify-content: flex-start;
  display: flex;
  height: 50%;
`;
const Date = styled.div `
  font-size: large;
`;
const Book = ({ match }) => {

    console.log(match.params.id);
    const [rbook, setRbook] = useState([]);
    const findBook = async () => {
        const res = await instance.get(`/bookread/${match.params.id}`);
        setRbook(res.data);
        console.log(res.data);
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
                </div>
    {/*            <div*/}
    {/*className="ck-content"*/}
    {/*dangerouslySetInnerHTML={{__html: rbook.subject}}*/}
    {/*/>*/}
                <CKEditor
                    s
                    editor={ ClassicEditor }
                    disabled={true}
                    config={ editorConfiguration }
                    data={rbook.subject}
                    onReady={ editor => {
                        // You can store the "editor" and use when it is needed.
                        console.log( 'Editor1 is ready to use!', editor );
                    } }
                />
            </BookDiv>

        </>
    );
};
export default Book;
