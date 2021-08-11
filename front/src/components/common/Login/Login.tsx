import React, {FC} from 'react';
import {GoogleLogin} from 'react-google-login';

const Login: FC = () => {

    const responseGoogle = (response: any) => {
        console.log(response);
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
