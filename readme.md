# zsh-gocloud-auto

zsh-gocloud-auto is a simple scrapper and generator that produces a gcloud autocompletion file for zsh from the public reference guide.

the original code for the auto completion was taken from littleq0903 script [here]( https://github.com/littleq0903/gcloud-zsh-completion) it is a bit out of date but worked great for the commands it had. I started filling in the missing commands and realised i am way to lazy to do that so just created this to generate it from google's awesome (and very nicely structured) [reference guide](https://cloud.google.com/sdk/gcloud/reference/).



## Todo
* strip out \n
* add flag to ignore cache and re-scrape
* flags not being displayed in auto complete
* add git diff as part of run to the _gcloud file to preview changes
* think of way to handle auto completion that should return the result of a command e.g organisations billings