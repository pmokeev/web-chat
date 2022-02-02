import React from 'react';

const InputText = (props: { type: string, placeholder: string, onChange: (target: string) => void, defaultValue: string }) => {
  return (
    <input
      type={props.type ? props.type : 'text'}
      placeholder={props.placeholder ? props.placeholder : ''}
      onChange={e => props.onChange(e.target.value)}
      value = {props.defaultValue}
    />
  )
}

export default InputText