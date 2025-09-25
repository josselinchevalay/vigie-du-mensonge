export function IndexArticles() {
    return (
        <>
            <div className="mx-auto w-full max-w-sm">
                <div className="flex flex-col items-center justify-center gap-4">
                    <h3>CLEMENT J'ATTENDS TOUJOURS LES CRITÈRES D'ACCEPTATION</h3>
                </div>
                <br/>
            </div>

            <div className="w-full flex flex-col sm:flex-row gap-4 items-center justify-between">
                <div className="flex flex-col">
                    <img src="/adso.jpg" alt="adso" className="w-100 h-100"/>
                    <h3>Le plus beau</h3>
                </div>

                <h3 className="text-xl">{"ADSO > HIPPIAS"}</h3>

                <div className="flex flex-col">
                    <img src="/hippias.PNG" alt="hippias" className="w-100 h-100"/>
                    <h3>C'est quoi ce regard ??</h3>
                </div>
            </div>
        </>
    );
}

// function IndexArticleNavButton(props: { article: { id: string } }) {
//     return (
//         <Link
//             to="/redactor/article-form"
//             search={{articleID: props.article.id}}
//             className="inline-flex items-center"
//         >
//             <Eye/>
//         </Link>
//     );
// }