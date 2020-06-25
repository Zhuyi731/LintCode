var simplifyPath = function (path) {
  //   path = path.replace(/(\.{2,})/g, "..");
  path = path.replace(/(?<!\.)(\.\/)/g, "");
  path = path.replace(/(\/+)/g, "/");
  if (/^\//.test(path)) {
    path = path.substr(1, path.length);
  }

  if (/\/$/.test(path)) {
    path = path.substr(0, path.length - 1);
  }

  let pathList = path.split("/");
  let pathStack = [];

  pathList.forEach((el) => {
    if (el === "..") {
      if (pathStack.length) {
        pathStack.pop();
      }
    } else if (el === ".") {
    } else {
      pathStack.push(el);
    }
  });

  let result = "/" + pathStack.join("/");
  console.log(result);
  return result;
};

// simplifyPath('/home/') // /home
simplifyPath("/../"); // /
simplifyPath("/home//foo/"); // /home/foo
simplifyPath("/a/./b/../../c/"); // c
simplifyPath("/a/../../b/../c//.//"); //c
simplifyPath("/a//b////c/d//././/.."); //"/a/b/c"
simplifyPath("/..."); // /...
simplifyPath("/."); // /...
