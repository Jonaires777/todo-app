import React, { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement>{
    label: string;
}

const Input: React.FC<InputProps> = ({label, ...rest}) => {
    return(
        <div className="flex flex-row gap-2 text-[#ffffff]">
            <label htmlFor={label}>{label}</label>
            <input {...rest} className="rounded-lg"/>
        </div>
    )
}

export default Input