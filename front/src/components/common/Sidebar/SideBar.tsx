import React, {FC, useEffect, useState} from "react";
import styled from 'styled-components';
import {Link} from "react-router-dom";
import {instance} from "../../../axios";
import {IconContext} from "react-icons";
import {AiOutlineClose, AiOutlineMenu} from "react-icons/all";

import "./SideBar.css"


const Nav = styled.div`
  display: flex;
  justify-content: flex-start;
  align-content: center;
  height: 5rem;
  background-color: #8FB399;
`;

const SidebarNav = styled.div<{ sidebar: boolean }>`
  width: 20rem;
  height: 100vh;
  background-color: #8FB399;
  position: fixed;
  top: 0;
  left: ${({sidebar}) => (sidebar ? '0' : '-100%')};
`;

const NavIcon = styled(Link)`
  display: flex;
  margin-top: 1.5rem;
  justify-content: flex-start;
  align-content: center;
  font-size: 2rem;
  height: 5rem;
  margin-left: 2rem;
`;

const SidebarWrap = styled.div`
`;

const SidebarLink = styled(Link)`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 2rem;
  font-size: 3rem;
  padding: 2rem;
  text-decoration: none;
  font-family: Neodgm, serif;
  color: #ffffff;

  :active {
    color: black;
  }
`

const LogoLink = styled(Link)`
  align-items: center;
  display: flex;
  justify-content: center;
  font-size: 3rem;
  font-family: Neodgm, serif;
  color: #ffffff;
  text-align: center;
  margin-left: 30rem;
  
  :active {
    color: black;
  }
`;


const SideBar: FC = () => {

    const [sidebar, setSidebar] = useState(false);
    const showSidebar = () => setSidebar(!sidebar)
    const [units, setUnits] = useState<any[]>([]);

    const findUnits = async () => {
        await instance.get('/unitshosting').then(
            function (res: { data: []; }) {
                setUnits(res.data)
            })
            .catch(function (error: any) {
                    console.log(error)
                }
            );
    };

    useEffect(() => {
        findUnits().then(r => console.log(r));
    }, [])


    return (
        <IconContext.Provider value={{color: '#fff'}}>
            <Nav>
                <NavIcon to="#" onClick={showSidebar}>
                    <AiOutlineMenu/>
                </NavIcon>
                <LogoLink to="/">Sung.Blog</LogoLink>
            </Nav>
            <SidebarNav sidebar={sidebar}>
                <SidebarWrap>
                    <NavIcon to="#" onClick={showSidebar}>
                        <AiOutlineClose/>
                    </NavIcon>
                    <SidebarLink to={`/`} onClick={showSidebar}>
                        Home
                    </SidebarLink>
                    <SidebarLink to={`/all`} onClick={showSidebar}>
                        All
                    </SidebarLink>
                    {units.map(unit => (
                            <SidebarLink key={unit.id} to={`/unit/${unit.id}`} onClick={showSidebar}>
                                {unit.content_name}
                            </SidebarLink>
                        )
                    )}
                </SidebarWrap>
            </SidebarNav>
        </IconContext.Provider>
    );

};

export default SideBar
