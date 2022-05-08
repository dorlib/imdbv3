import classes from "./Layout.module.css";
import MainNavigation from "./MainNavigation";
import React, {useState} from "react";
import SearchBar from "./search-bar";

function Layout(props) {

    return (
        <div>
            <div>
                <MainNavigation/>
                <main className={classes.main}>{props.children}</main>
            </div>
        </div>
    );
}

export default Layout;
