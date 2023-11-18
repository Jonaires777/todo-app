import React, { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement>{
    label: string;
}

const Input: React.FC<InputProps> = ({label, ...rest}) => {
    return(
        <div className="flex flex-row gap-2">
            <label className="text-[#ffffff]" htmlFor={label}>{label}</label>
            <input {...rest} className="rounded-lg p-1"/>
        </div>
    )
}

export default Input