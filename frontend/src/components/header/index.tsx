import { FC, memo } from 'react';
import logo from '../../assets/images/logo.png';
import close from '../../assets/images/close.png';
import './index.css';

const Header: FC<WailsProps & { handelClose: Function}> = memo(({ DataTitle, DataIcon, handelClose }: WailsProps & { handelClose: Function}) => {
    console.log("DataIco:",DataIcon)
    return (
        <div className='header'>
            <div className='left'>
                <img className='logo' src={DataIcon === "default" ? logo : DataIcon} alt="logo" />
                <span className='lefetitle'>{DataTitle}</span>
            </div>
            <div className='right'>
                <img className='close' onClick={() => handelClose()} src={close} alt="关闭" />
            </div>
        </div>
    )
})

export default Header;
