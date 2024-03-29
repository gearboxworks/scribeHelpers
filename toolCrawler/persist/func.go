package persist

import (
	"encoding/json"
	"fmt"
	"github.com/gearboxworks/go-status/only"
	"github.com/jtacoma/uritemplates"
	"github.com/mattn/go-sqlite3"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/config"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/global"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/pages"
	"github.com/gearboxworks/scribeHelpers/toolCrawler/util"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var template *uritemplates.UriTemplate

func noop(i ...interface{}) interface{} { return i }

func init() {
	template, _ = uritemplates.Parse(JsonFileTemplate)
}

func IsSqlUniqueError(err error) (is bool) {
	for range only.Once {
		err, ok := err.(sqlite3.Error)
		if !ok {
			break
		}
		if err.Code != sqlite3.ErrConstraint {
			break
		}
		if err.ExtendedCode != sqlite3.ErrConstraintUnique {
			break
		}
		is = true
	}
	return is
}

func getRootUrl(inurl global.Url) (outurl global.UrlPath, err error) {
	for range only.Once {
		inptr, err := parseUrl(inurl)
		if err != nil {
			break
		}
		s := inptr.Scheme
		d := inptr.Hostname()
		p := inptr.Port()
		if p != "" && p != "80" {
			p = fmt.Sprintf(":%s", p)
		}
		outurl = fmt.Sprintf("%s://%s%s", s, d, p)
	}
	return outurl, err
}

func parseUrl(u global.Url) (uu *url.URL, err error) {
	uu, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("unable to parse URL '%s'", u)
		logrus.Error(err)
	}
	return uu, err
}

func getPathFromUrl(u global.Url) (p global.UrlPath, err error) {
	for range only.Once {
		uptr, err := parseUrl(u)
		if err != nil {
			break
		}
		p = getParsedPath(*uptr)
	}
	return p, err
}

func getParsedPath(uptr url.URL) (p global.UrlPath) {
	for range only.Once {
		p = uptr.Path
		if uptr.RawQuery != "" {
			p = fmt.Sprintf("%s?%s", p, uptr.RawQuery)
		}
	}
	return p
}

func HasQueuedUrls(cfg *config.Config) (found bool) {
	for range only.Once {
		qsd := GetQueuedSubdir(cfg)
		f, err := os.Open(qsd)
		if err != nil {
			logrus.Fatalf("unable to open queued directory '%s': '%s'",
				qsd,
				err.Error(),
			)
			break
		}
		// Assumes no more than a few non-directories
		files, err := f.Readdir(10)
		if err != nil {
			logrus.Fatalf("unable to read 10 directory entries in queued directory '%s': '%s'",
				qsd,
				err.Error(),
			)
			break
		}
		for _, f := range files {
			if !f.IsDir() {
				continue
			}
			if len(f.Name()) != 2 {
				break
			}
			found = true
			break
		}
	}
	return found
}

func GetQueuedUrls(cfg *config.Config) (global.Urls, error) {
	qsd := GetQueuedSubdir(cfg)
	queued := make(global.Urls, 0)
	err := filepath.Walk(qsd, func(fp string, fi os.FileInfo, err error) error {
		for range only.Once {
			if err != nil {
				break
			}
			if fi.Name() == ".DS_Store" {
				break
			}
			if fi.IsDir() {
				break
			}
			c, err := ioutil.ReadFile(fp)
			if err != nil {
				err = fmt.Errorf("cannot read file '%s': %s", fp, err.Error())
			}
			queued = append(
				queued,
				strings.TrimSpace(global.Url(c)),
			)
		}
		return err
	})
	if err != nil {
		err = fmt.Errorf("cannot walk directory '%s': %s", qsd, err.Error())
	}
	return queued, err
}

//func QueueUrl(cfg *config.Config, url global.Url) (found bool) {
//	for range only.Once {
//		fp, err := GetSubdirFilepath(cfg, QueuedDir, url)
//		if err != nil {
//			break
//		}
//		if WriteFile(fp, []byte(url), CannotExist) != nil {
//			break
//		}
//		found = true
//	}
//	return found
//}

func IndexedPage(cfg *config.Config, page *pages.Page) (err error) {
	return Persist(cfg, IndexedDir, page)
}
func ErroredPage(cfg *config.Config, page *pages.Page) (err error) {
	return Persist(cfg, ErroredDir, page)
}
func Persist(cfg *config.Config, subdir global.Dir, page *pages.Page) (err error) {
	for range only.Once {
		var b []byte
		b, err = json.Marshal(page)
		if err != nil {
			err = fmt.Errorf("unable to marshal page '%s': %s", page.Url, err.Error())
			break
		}
		fp, err := GetSubdirFilepath(cfg, subdir, page.Url)
		if err != nil {
			break
		}
		err = WriteFile(fp, b, CanExist)
		if err != nil {
			err = fmt.Errorf("%s for URL '%s'", err.Error(), page.Url)
			break
		}
	}
	return err
}

func WriteFile(fp global.Filepath, content []byte, exists Existence) (err error) {
	for range only.Once {
		switch exists {
		case CannotExist:
			if util.FileExists(fp) {
				err = fmt.Errorf("file '%s'already exists", fp)
			}
		case MustExist:
			if !util.FileExists(fp) {
				err = fmt.Errorf("file '%s' does not exists", fp)
			}
		case CanExist:
			// Do nothing
		}
		if err != nil {
			break
		}
		err = os.MkdirAll(filepath.Dir(fp), os.ModePerm)
		if err != nil {
			logrus.Errorf("unable to make directory '%s'", fp)
			break
		}
		err = ioutil.WriteFile(fp, content, os.ModePerm)
		if err != nil {
			err = fmt.Errorf("unable to write file '%s': %s", fp, err.Error())
			break
		}
	}
	return err
}

func GetUrlFilename(url global.Url) (fn global.Filename, err error) {
	for range only.Once {
		if !pages.IsIndexable(url) {
			err = fmt.Errorf("the URL '%s' is not an indexable URL", url)
			break
		}
		h := pages.NewHash(url)
		fn, err = template.Expand(map[string]interface{}{
			"hash": h,
		})
		if err != nil {
			logrus.Errorf("unable to expand template '%s' with hash='%s'",
				JsonFileTemplate,
				h,
			)
			break
		}
	}
	return fn, err
}

func GetQueuedSubdir(cfg *config.Config) (d global.Dir) {
	return GetSubdir(cfg, QueuedDir)
}

func GetIndexedSubdir(cfg *config.Config) (d global.Dir) {
	return GetSubdir(cfg, IndexedDir)
}

func GetErroredSubdir(cfg *config.Config) (d global.Dir) {
	return GetSubdir(cfg, ErroredDir)
}

func GetSubdir(cfg *config.Config, subdir global.Dir) (d global.Dir) {
	return fmt.Sprintf("%s%c%s",
		cfg.DataDir,
		os.PathSeparator,
		subdir,
	)
}

func GetSubdirFilepath(cfg *config.Config, subdir global.Dir, url global.Url) (fp global.Filepath, err error) {
	for range only.Once {
		var fn global.Filename
		fn, err = GetUrlFilename(url)
		if err != nil {
			break
		}
		fp = fmt.Sprintf("%s%c%s",
			GetSubdir(cfg, subdir),
			os.PathSeparator,
			fn,
		)
	}
	return fp, err
}

func GetDbFilepath(cfg *config.Config) global.Filepath {
	return fmt.Sprintf("%s%c%s",
		cfg.DataDir,
		os.PathSeparator,
		SqliteDbFilename,
	)
}
