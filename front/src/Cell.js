import React from 'react'
import './Cell.css'

class Cell extends React.Component {
    renderCell = (isChecked, byPlayer) => {
        if (isChecked) {
            switch (byPlayer) {
                case 1:
                    return 'cross.png';
                case 2:
                    return 'circle.png';
                default:
                    return '';
            }
        }
        return '';
    }

    render() {
        let content = this.renderCell(this.props.isChecked, this.props.byPlayer);

        return (
            <div className={"cell"}>
                <img src={content} alt=""/>
            </div>
        )
    }
}

export default Cell