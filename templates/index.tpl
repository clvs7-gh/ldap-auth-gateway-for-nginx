<html lang="ja">
    <head>
        <meta charset="utf-8" />
        <meta name="robots" content="noindex,nofollow,noarchive" />
        <link rel="stylesheet" type="text/css" href="./assets/main.css" />
        <title>SECRET AREA | LOGIN</title>
    </head>
    <body>
        <div id="background__container">
            <span id="background__container__text">S</span>
        </div>
        <div id="container">
            <div id="container__login__box">
                <h2 id="container__login__box__title">SECRET AREA</h2>
                <h3 class="container__login__box__description">Welcome</h3>
                <h3 class="container__login__box__message">{{.Message}}</h3>
                <form class="container__login__box__input" method="POST" action="login">
                        <input type="text" name="username" placeholder="Username" class="container__login__box__input--username" required />
                        <input type="password" name="password" placeholder="Password" class="container__login__box__input--password" required />                    
                        <input type="hidden" name="redirect_to" value="{{.RedirectTo}}" />                    
                        <input type="submit" value="Login" class="container__login__box__input--submit" />                    
                </form>
                <div id="container__login__box__forget">
                    <a href="#" onclick="onForgotPassword(event)">Forgot password</a>
                </div>
            </div>
        </div>
        <script type="text/javascript" src="./assets/main.js"></script>
    </body>
</html>
