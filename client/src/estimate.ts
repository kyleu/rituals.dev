namespace estimate {
  interface Detail extends rituals.Session {
    choices: string[];
    options: object;
  }

  export interface StoryStatusChange {
    storyID: string;
    status: string;
    finalVote: string;
  }

  interface SessionJoined extends rituals.SessionJoined {
    session: Detail;
    stories: story.Story[];
    votes: vote.Vote[];
  }

  class Cache {
    activeStory?: string;

    detail?: Detail;

    stories: story.Story[] = [];
    votes: vote.Vote[] = [];

    public activeVotes(): vote.Vote[] {
      if (this.activeStory === undefined) {
        return [];
      }
      return this.votes.filter(x => x.storyID == this.activeStory);
    }
  }

  export const cache = new Cache();

  export function onEstimateMessage(cmd: string, param: any) {
    switch (cmd) {
      case command.server.error:
        rituals.onError(services.estimate, param as string);
        break;
      case command.server.sessionJoined:
        let sj = param as SessionJoined;
        rituals.onSessionJoin(sj);
        setEstimateDetail(sj.session);
        story.setStories(sj.stories);
        vote.setVotes(sj.votes);
        break;
      case command.server.sessionUpdate:
        setEstimateDetail(param as Detail);
        break;
      case command.server.storyUpdate:
        onStoryUpdate(param as story.Story);
        break;
      case command.server.storyRemove:
        onStoryRemove(param as string);
        break;
      case command.server.storyStatusChange:
        story.onStoryStatusChange(param as StoryStatusChange);
        break;
      case command.server.voteUpdate:
        vote.onVoteUpdate(param as vote.Vote);
        break;
      default:
        console.warn("unhandled command [" + cmd + "] for estimate");
    }
  }

  function setEstimateDetail(detail: Detail) {
    cache.detail = detail;
    util.setValue("#model-choices-input", detail.choices.join(", "));
    story.viewActiveStory();
    rituals.setDetail(detail);
  }

  export function onSubmitEstimateSession() {
    const title = util.req<HTMLInputElement>("#model-title-input").value;
    const choices = util.req<HTMLInputElement>("#model-choices-input").value;
    const msg = {
      svc: services.estimate,
      cmd: command.client.updateSession,
      param: {
        title: title,
        choices: choices,
      },
    };
    socket.send(msg);
  }

  export function onStoryUpdate(s: story.Story) {
    const x = preUpdate(s.id)
    x.push(s);
    if(s.id === estimate.cache.activeStory) {
      util.setText("#story-title", s.title);
    }
    story.setStories(x);
  }

  export function onStoryRemove(id: string) {
    const x = preUpdate(id)
    story.setStories(x);
    if(id === estimate.cache.activeStory) {
      UIkit.modal("#modal-story").hide();
    }
    UIkit.notification("story has been deleted", {status: "success", pos: "top-right"});
  }

  function preUpdate(id: string) {
    return estimate.cache.stories.filter((p) => p.id !== id);
  }
}
