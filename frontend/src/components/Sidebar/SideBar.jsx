import React, { useEffect, useState } from "react";
import styled from 'styled-components';
import { Link } from "react-router-dom";
import { instance } from "../../axios";
import { IconContext } from "react-icons";
import { AiOutlineClose, AiOutlineMenu } from "react-icons/all";
import { GoogleLogin } from "react-google-login";

const Nav = styled.div `
  display: flex;
  justify-content: flex-start;
  align-content: center;
  height: 5rem;
  background-color: #8FB399;
`;
const SidebarNav = styled.div `
  width: 16rem;
  height: 100vh;
  background-color: #8FB399;
  position: fixed;
  top: 0;
  left: ${({ sidebar }) => (sidebar ? '0' : '-100%')};
`;
const NavIcon = styled(Link) `
  display: flex;
  margin-top: 1.5rem;
  justify-content: flex-start;
  align-content: center;
  font-size: 2rem;
  height: 5rem;
  margin-left: 2rem;
`;
const SidebarWrap = styled.div `
`;
const SidebarLink = styled(Link) `
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 2rem;
  font-size: 2rem;
  padding: 2rem;
  text-decoration: none;
  font-family: Neodgm, serif;
  color: #ffffff;

  :active {
    color: black;
  }
`;
const LogoLink = styled(Link) `
  align-items: center;
  display: flex;
  justify-content: center;
  font-size: 200%;
  font-family: Neodgm, serif;
  color: #ffffff;
  text-align: center;
  margin-left: 35%;

  :active {
    color: black;
  }
`;
const MakeLink = styled(Link) `
  font-family: Neodgm, serif;
  color: #ffffff;
`;

const SideBar = () => {
    const [sidebar, setSidebar] = useState(false);
    const [login, setLogin] = useState(!!sessionStorage.getItem('id'));
    const showSidebar = () => setSidebar(!sidebar);
    const [units, setUnits] = useState([]);

    const responseGoogle = async (response) => {
        const res = await instance.post('/login', {
            num: response.Ts.mS
        });
        if (res.data == null) {
            console.log("error");
        }
        else {
            window.sessionStorage.setItem('id', res.data[0].id);
            setLogin(true);
        }
    };

    const findUnits = async () => {
        await instance.get('/unitshosting').then(function (res) {
            setUnits(res.data);
        })
            .catch(function (error) {
                console.log(error);
            });
    };

    const logout = () => {
        localStorage.removeItem('id');
        setLogin(false);
    };

    useEffect(() => {
        findUnits();
    }, []);

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
                    <div>
                        {!login ? <GoogleLogin
                            clientId={"940522265963-gqbtd1jmbqtsueje1hhfqved273412i7.apps.googleusercontent.com"}
                            buttonText="Google"
                            render={renderProps => (
                                <a
                                    className="nes-btn"
                                    href="#"
                                    onClick={renderProps.onClick}
                                    style={{marginLeft: "10%", fontFamily: "Neodgm", fontSize: "xx-large"}}
                                >
                                    Login
                                </a>
                            )}
                            onSuccess={responseGoogle}
                            onFailure={responseGoogle}
                        /> : <>
                            <MakeLink to={'/bookwrite'}>
                                <
                                    button
                                    type="button"
                                    className="nes-btn is-success"
                                    style={{marginLeft: "10%", fontFamily: "Neodgm", fontSize: "x-large"}}
                                    onClick={showSidebar}
                                >
                                    write
                                </button>
                            </MakeLink>
                            <
                                button
                                type="button"
                                className="nes-btn is-error"
                                style={{marginLeft: "10%", fontFamily: "Neodgm", fontSize: "x-large"}}
                                onClick={logout}
                            >
                                logout
                            </button>
                        </>}
                    </div>
                </SidebarWrap>
            </SidebarNav>
        </IconContext.Provider>
    );
};

export default SideBar

