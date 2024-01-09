import './index.css';
import logo from '../../assets/images/logo.png';

export type HeaderProps = {
    title?: string;
}

function Header(props: HeaderProps) {
    const { title = "通知" } = props
    return (
        <div className='header'>
            <div className='left'>
                <img className='logo' src={logo} alt="" />
                {title}
            </div>
            <div className='right'></div>
        </div>
    )
}

export default Header;
