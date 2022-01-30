import React, {SyntheticEvent, useEffect, useState} from 'react';
import {Redirect} from "react-router-dom";

const Profile = (props: { isJWTCorrect: boolean }) => {
  let formSubmit;

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [oldPassword, setOldPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [submitStatusCode, setSubmitStatusCode] = useState(0);
  const [isHidden, setIsHidden] = useState(true);

  useEffect(() => {
    (
      async () => {
        const response = await fetch('http://localhost:8000/api/auth/profile', {
          method: 'GET',
          credentials: 'include',
        });

        const content = response.json();
        content.then(data => {
          setName(data["name"]);
          setEmail(data["email"])
        });
      }
    )();
  });

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    const response = await fetch('http://localhost:8000/api/auth/change-password', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      credentials: 'include',
      body: JSON.stringify({
        email,
        oldPassword,
        newPassword,
      })
    });

    setSubmitStatusCode(response.status)
    if (response.status === 200 || response.status === 400) {
      setIsHidden(true);
    }
  }

  if (!props.isJWTCorrect) {
    return <Redirect to="/login"/>
  }

  if (!isHidden) {
    formSubmit = (
      <form onSubmit={submit}>
        <div className="form-floating">
          <input type="password" className="form-control" id="floatingPassword" placeholder="Old password"
                 onChange={e => setOldPassword(e.target.value)}
          />
          <label htmlFor="floatingPassword">Old password</label>
        </div>
        <div className="form-floating">
          <input type="password" className="form-control" id="floatingPassword" placeholder="New password"
                 onChange={e => setNewPassword(e.target.value)}
          />
          <label htmlFor="floatingPassword">New password</label>
        </div>
        <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
      </form>
    )
  } else if (submitStatusCode === 200) {
    formSubmit = (
      <div className="helloMsg">Password has been successfully changed </div>
    )
  } else if (submitStatusCode === 400) {
    formSubmit = (
      <div className="ErrorMsg">Incorrect old password</div>
    )
  }

  return (
    <div>
      <div className="helloMsg">
        Hello {name}! Your email - {email}
      </div>

      <button className="w-100 btn btn-lg btn-primary" onClick={() => setIsHidden(false)}>Change password</button>
      {formSubmit}
    </div>
  );
};

export default Profile;