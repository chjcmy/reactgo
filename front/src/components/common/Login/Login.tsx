import React, {FC} from 'react';
import {GoogleLogin} from 'react-google-login';
import {instance} from "../../../axios";

const Login: FC = () => {

    const responseGoogle = async (response: any) => {
        console.log(response.Ts.mS);
        const res = await instance.post('/login',{
            num: response.Ts.mS
        })
        console.log(res.data[0].name)
        window.localStorage.setItem("name", JSON.stringify(res.data[0].name));
    }

    return (
        <>
            <GoogleLogin
                clientId={"940522265963-gqbtd1jmbqtsueje1hhfqved273412i7.apps.googleusercontent.com"}
                buttonText="Google"
                render={renderProps => (
                    <a
                        className="nes-btn"
                        href="#"
                        onClick={renderProps.onClick}
                        style={{marginLeft: "10%", fontFamily: "Neodgm", fontSize: "x-large"}}
                    >
                        Login
                    </a>
                )}
                onSuccess={responseGoogle}
                onFailure={responseGoogle}
            />
        </>
    );
};

export default Login;
